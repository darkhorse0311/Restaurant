const restaurantData = require('../restaurantData.json');

// console.log(restaurantData.length);
// console.log(Object.keys(restaurantData[0])[0]);
const allItems = [];

restaurantData.forEach((res, index) => {
  res[Object.keys(res)[0]].forEach((item, i) => {
    const { name, protein, carbs, fats, calories } = item;
    if( !name, !protein, !carbs, !fats, !calories ) {
      return;
    }
    if (name === 'Item') {
      return;
    }
    
    const newItem = {
      name: item.name,
      type: item.type,
      protein: Number(item.protein),
      carbs: Number(item.carb),
      fats: Number(item.fat),
      calories: Number(item.calories),
      calsperpro: Number(item.calperpro),
      sodium: Number(item.sodium),
      restaurant_id: index,
    };
    allItems.push(newItem)
  });
});


exports.seed = async function(knex, Promise) {
  // Deletes ALL existing entries

  for (const item of allItems) {
    await knex('items').insert([item]);
  }

  // return knex('items').insert(fith);
};
