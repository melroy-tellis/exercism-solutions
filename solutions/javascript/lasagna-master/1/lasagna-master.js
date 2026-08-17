/// <reference path="./global.d.ts" />
// @ts-check

/**
 * Implement the functions needed to solve the exercise here.
 * Do not forget to export them so they are available for the
 * tests. Here an example of the syntax as reminder:
 *
 * export function yourFunction(...) {
 *   ...
 * }
 */
export function cookingStatus(timer) {
  if (timer === undefined) {
    return 'You forgot to set the timer.'
  }
  if (timer === 0) {
    return 'Lasagna is done.'
  }
  return 'Not done, please wait.'
}

export function preparationTime(layers, avgTimePerLayer=2) {
  return avgTimePerLayer*layers.length;
}

export function quantities(layers) {
  let quant = {
    "noodles": 0,
    "sauce": 0
  };
  let quantitiesByLayer = {
    "noodles": 50,
    "sauce": 0.2
  };
  for (let layer of layers) {
    if (!Object.hasOwn(quantitiesByLayer, layer)) {
      continue;
    }
    quant[layer] += quantitiesByLayer[layer];
  }
  return quant;
}

export function addSecretIngredient(friendsList, myList) {
  myList.push(friendsList[friendsList.length-1]);
}

export function scaleRecipe(recipe, portions) {
  let scalingFactor = portions/2;
  let scaledRecipe = {...recipe};
  for (let ingredient in recipe) {
    scaledRecipe[ingredient] = recipe[ingredient] * scalingFactor;
  }
  return scaledRecipe;
  
}
