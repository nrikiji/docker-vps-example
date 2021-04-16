import axios from "../node_modules/axios/index";

function main() {
  axios.get("https://kimagurenews.xyz/").then((x) => console.log(x.data));
}

main();
