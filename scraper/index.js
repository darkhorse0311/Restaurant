const axios = require('axios');
const cheerio = require('cheerio');
const fs = require('fs');

const resturant = {};

const url = 'http://fastfoodmacros.com/food.asp?n=McDonald%27s+Macros&r=1';

const


axios.get(url)
    .then(res => {
        const $ = cheerio.load(res.data)

        const rows = $('tbody').children();

        let menuItems = [];
        rows.each((index, element) => {
            let item = {
                name: null,
                type: null,
                protein: null,
                fat: null,
                carb: null,
                calories: null,
                calperpro: null,
                sodium: null
            };

            $(element).children().each((i, el) => {
                let section = {};
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
                        item.protein = $(el).first().text();
                        break;
                    case 4 : 
                        item.fat = $(el).first().text();
                        break;
                    case 5 : 
                        item.carb = $(el).first().text();
                        break;
                    case 6 : 
                        item.calories = $(el).first().text();
                        break;
                    case 7 : 
                        item.calperpro = $(el).first().text();
                        break;
                    case 8 : 
                        item.sodium = $(el).first().text();
                        break;
                    default: ;
                }
                // console.log(item.name);
            })
            menuItems.push(item);
        });
        // console.log('menuItems', menuItems);
        // console.log(menuItems.length);
        const resturants = {
            mcdonalds: {
                items: menuItems
            }
        }

        // console.log(menuItems);
        fs.writeFileSync('./mcdonaldsData.json', JSON.stringify(resturants))
    })
    .catch(err => console.log(err));
