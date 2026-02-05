// Создаём обещание: ждём 2 секунды
const pizzaPromise = new Promise((resolve, reject) => {
  setTimeout(() => {
    if (Math.random() > 0.5) {
      resolve("Пицца готова! 🍕"); // Успех
    } else {
      reject("Пиццы нет! 😞"); // Ошибка
    }
  }, 2000);
});

pizzaPromise
  .then((pizza) => console.log(pizza)) // Если успех
  .catch((error) => console.error(error)) // Если ошибка
  .finally(() => console.log("Done."));

let x = 5;

function makeX() {
  console.log(x);

  function inc() {
    x++;
  }

  msg = `X is ${x}`;
  function log() {
    console.log(msg);
  }

  return [inc, log];
}

let [inc, log] = makeX();
inc();
inc();
console.log(x);
log();
