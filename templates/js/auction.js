function ShowHideDiv() {
  var auctionHouse = document.getElementById("playerOnTeams");

  //var numberOfPlayers2v2 = document.getElementById("teamPicker2v2");
  //var numberOfPlayers2v2 = document.getElementById("teamPicker2v2");
  numberOfPlayers2v2.style.display = playersOnTeams.value == "2v2" ? "block" : "none";
  numberOfPlayers3v3.style.display = playersOnTeams.value == "3v3" ? "block" : "none";
  numberOfPlayers4v4.style.display = playersOnTeams.value == "4v4" ? "block" : "none";
  numberOfPlayers5v5.style.display = playersOnTeams.value == "5v5" ? "block" : "none";
}

// Check if clear teams button is clicked
document.getElementById("clearTeams2v2").addEventListener("click", reload);
document.getElementById("clearTeams3v3").addEventListener("click", reload);
document.getElementById("clearTeams4v4").addEventListener("click", reload);
document.getElementById("clearTeams5v5").addEventListener("click", reload);

//function for reloading the page
function reload() {
  window.location.reload();
}

function ShowRandomTeams2v2() {
  document.getElementById("team12v2").style.display = "block";
  document.getElementById("vs2v2").style.display = "block";
  document.getElementById("team22v2").style.display = "block";
}
