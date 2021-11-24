import axios from 'axios';

export default class todo {
  static fetchAll() {
    axios.get('https://jsonplaceholder.typicode.com/todos/1')
      .then((response) => console.log(response));
  }
}
