import axios from "../node_modules/axios/index";

function main() {
  axios.get("https://backend.example.com/").then((x) => console.log(x.data));
}

main();
