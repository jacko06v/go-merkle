const fs = require("fs");

fs.readFile("renting_int_ids.json", "utf8", function (err, data) {
  if (err) {
    console.error(err);
    return;
  }

  const jsonString = JSON.stringify(
    JSON.parse(data),
    (_, value) => {
      if (typeof value === "number") {
        return value.toString();
      }
      return value;
    },
    2
  );

  fs.writeFile("output_string.json", jsonString, "utf8", function (err) {
    if (err) {
      console.error(err);
      return;
    }
    console.log("Conversione completata!");
  });
});
