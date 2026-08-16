//
// This is only a SKELETON file for the 'Largest Series Product' exercise. It's been provided as a
// convenience to get you started writing code faster.
//

export const largestProduct = (input, span) => {
  if (input.length < span) {
    throw new Error("span must not exceed string length");
  }
  if (span < 0) {
    throw new Error("span must not be negative");
  }
  if (!/^\d+$/.test(input)) {
    throw new Error("digits input must only contain digits");
  }
  let maxProduct = 0;
  for (let i = 0; i <= input.length-span; i++) {
    let product = 1;
    for (let j = 0; j < span; j++) {
      product *= Number.parseInt(input[i+j]); 
    }
    maxProduct = Math.max(product, maxProduct);
  }
  return maxProduct;
};
