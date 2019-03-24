const axios = require('axios');
const cheerio = require('cheerio');
const fs = require('fs');

const resturant = {};

// const url = 'http://fastfoodmacros.com/food.asp?n=McDonald%27s+Macros&r=1';
// const url = 'http://fastfoodmacros.com/food.asp?n=Jack+in+the+Box+Macros&r=24';
const url = 'http://fastfoodmacros.com#';

const getResturantUrls = (page) => {
    const $ = cheerio.load(page)

    const resturantUrls = [];

    const resturantList = $('.pushy-submenu > ul > li');
    
    resturantList.each((index, element) => {
        const pre = 'http://fastfoodmacros.com/';
        const link = $(element).children().first().attr('href');

        if (link.includes('food.asp')){
            resturantUrls.push(`${pre}${link}`);
        }

    })
    return resturantUrls;
}

const getResturantInfo = (page) => {
    const $ = cheerio.load(page)

    title = $('head > title').text().split(' ');
    resName = title.slice(0, title.length - 3).join('');

    const rows = $('tbody').children();
    let menuItems = [];
    rows.each((index, element) => {
        let item = {
            name: null,
            type: null,
            protein: null,
            fats: null,
            carbs: null,
            calories: null,
            calperpro: null,
            sodium: null
        };

        $(element).children().each((i, el) => {
            switch(i) {
                case 1 :
                    // section = {type: $(el).first().text()};
                    item.name = $(el).children().first().text();
                    break;
                case 2 :
                    // section = {type: $(el).first().text()};
                    item.type = $(el).first().text();
                    break;
                case 3 : 
                    item.protein = Number($(el).first().text());
                    break;
                case 4 : 
                    item.fats = Number($(el).first().text());
                    break;
                case 5 : 
                    item.carbs = Number($(el).first().text());
                    break;
                case 6 : 
                    item.calories = Number($(el).first().text());
                    break;
                case 7 : 
                    item.calperpro = Number($(el).first().text());
                    break;
                case 8 : 
                    item.sodium = Number($(el).first().text());
                    break;
                default: ;
            }
        })

        if (item.name != null && item.name != "Item"){
            menuItems.push(item);
        }
    });

    return {name: resName, items: menuItems};
}

const start = async () => {
    let resturants = [];
    const homePage = await axios.get(url);
    const resturantUrls = getResturantUrls(homePage.data);

    for (const link of resturantUrls) {
        const resPage = await axios.get(link);
        const restInfo = getResturantInfo(resPage.data);
        resturants.push(restInfo);
    }
    fs.writeFileSync('./restaurantData.json', JSON.stringify(resturants))
}

start();