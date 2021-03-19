function ShowHideDiv() {
  var realmWanted = document.getElementById("realm");
  var auctionHouse = document.getElementById("auctionHouse");

  //var numberOfPlayers2v2 = document.getElementById("teamPicker2v2");
  //var numberOfPlayers2v2 = document.getElementById("teamPicker2v2");
  auctionHouse.style.display = realmWanted.value !== "pick" ? "block" : "none";
}

// Check if clear teams button is clicked
document.getElementById("clearRealm").addEventListener("click", reload);

//function for reloading the page
function reload() {
  window.location.reload();
}
