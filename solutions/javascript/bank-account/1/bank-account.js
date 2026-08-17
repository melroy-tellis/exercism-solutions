//
// This is only a SKELETON file for the 'Bank Account' exercise. It's been provided as a
// convenience to get you started writing code faster.
//

export class BankAccount {
  #balance;
  #state;
  
  constructor() {
    this.#balance = 0;
    // 0 = initial, 1 = open, 2 = closed
    this.#state = 0;
  }

  open() {
    if (this.#state != 1) {
      this.#state = 1;
      this.#balance = 0;
    } else {
      throw new ValueError();
    }
  }

  close() {
    if (this.#state == 1) {
      this.#state = 2;
    } else {
      throw new ValueError();
    }
  }

  deposit(amount) {
    if (this.#state == 1) {
      if (amount < 0) {
        throw new ValueError();
      }
      this.#balance += amount;
    } else {
      throw new ValueError();
    }
    
  }

  withdraw(amount) {
    if (this.#state == 1) {
      if (amount<0 || amount>this.#balance) {
        throw new ValueError();
      }
      this.#balance -= amount;
    } else {
      throw new ValueError();
    }
  }

  get balance() {
    if (this.#state == 1) {
      return this.#balance;
    } else {
      throw new ValueError();
    }
  }
}

export class ValueError extends Error {
  constructor() {
    super('Bank account error');
  }
}
