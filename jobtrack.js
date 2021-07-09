function currentDate() {
    var options = { month: "long", day: "numeric", year: "numeric" };
    document.write(new Date().toLocaleDateString("en-CA", options));
}
