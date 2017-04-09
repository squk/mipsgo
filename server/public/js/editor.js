$(document).ready(function() {
  // Init the top cell content
  var topContent = "";

  for(i=0; i<16; i++) {
    topContent += (0+i.toString(16)).slice(-2)+" ";
  }
  $("#topE").text(topContent);

  var h = "";
  var b, c;


  $("#middleE")[0].oninput = refreshMemory;

  // (4 bits per hex digit * 32 bits per word) * (2 << D) words
  for(i=0; i<(2<<12); i++) {
    memory += "00 ";
  }
  setMemory();

  $("#clear_memory").click(function() {
    for(i=0; i<(2<<12); i++) {
      memory += "00 ";
    }
    set_memrange(0);
  });
});

function refreshMemory() {
  var middle = $("#middleE")[0];

  // On input, store the length of clean hex before the textarea caret in b
  b = middle.value
    .substr(0,middle.selectionStart)
    .replace(/[^0-9A-F]/ig,"")
    .replace(/(..)/g,"$1 ")
    .length;

  // Clean the textarea value
  $(middle).val(middle.value
    .replace(/[^0-9A-F]/ig,"")
    .replace(/(..)/g,"$1 ")
    .replace(/ $/,"")
    .toUpperCase()
  );

  // Reset h
  h="";

  // Loop on textarea lines
  for(i=0;i<middle.value.length/48;i++) {
    // Add line number to h
    h += (1E7+(16*i).toString(16)).slice(-8)+" ";
  }

  // Write h on the left column
  $("#leftE").text(h);

  // Reset h
  h="";

  // Loop on the hex values
  for(i=0;i<middle.value.length;i+=3) {
    // Convert them in numbers
    c = parseInt(middle.value.substr(i,2),16);

    // Convert in chars (if the charCode is in [64-126] (maybe more later)) or ".".
    h = 63<c && 127>c ? h+String.fromCharCode(c) : h+".";
  }

  // Write h in the right column (with line breaks every 16 chars)
  $("#rightE").text(h.replace(/(.{16})/g,"$1 "));

  // If the caret position is after a space or a line break, place it at
  // the previous index so we can use backspace to erase hex code
  if(middle.value[b] == " ") {
    b--;
  }

  // Put the textarea caret at the right place
  middle.setSelectionRange(b,b)
}
