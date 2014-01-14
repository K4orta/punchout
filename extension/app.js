var btnMap = {};
var actions = [];
var punchServer = "http://localhost:8081";

var update = function() {
  actions = [];
  $('.wfa-clock-punchbutton').each(function() {
    if ($(this).css('display') !== 'none') {
      actions.push($(this).data('punch-action'));
    }
  });
  // DEBUG REMOVE
  actions = ['ClockInForDay','bleh']
  postStatus(actions);
};
var timer;

var punchAction = function(option) {
  // ClockInForDay
  // ClockInFromLunch
  // ClockOutForLunch
  // ClockOutForDay
  if(btnMap[option.data] == null) {
    console.log('Unknown command: ' + option.data);
    return;
  }
  btnMap[actions[0]].click();
  console.log(option.data);
};

var buildMap = function() {
  $('.wfa-clock-punchbutton').each(function() { 
    btnMap[$(this).data('punch-action')] = $(this);
  });
};

var source;
if(window.EventSource) {
  source = new window.EventSource(punchServer + '/commands', {withCredentials: true}) 
  source.addEventListener('punchEvent', punchAction);
}

var postStatus = function (actions) {
  // console.log(actions);
  if (actions.length === 1) {
    $.get(punchServer + "/update/" + "clockedOut");
  } else if (action.length == 2) {
    $.get(punchServer + "/update/" + "clockedIn");
  } else {
   $.get(punchServer + "/update/" + "clockedIn"); 
  }
};

$(document).ready(function() {
  $('body').css('background-color', '#999');
  console.log('running btn scrape');
  buildMap();
  timer = setInterval(function() {update()}, 3000);
  // punchAction('ClockOutForDay');
});