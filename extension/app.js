var btnMap = {}
var actions = [];

var update = function() {
  actions = [];
  $('.wfa-clock-punchbutton').each(function() {
    if ($(this).css('display') !== 'none') {
      actions.push($(this).data('punch-action'));
    }
  });
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
  btnMap[option.data].click();
};

var buildMap = function() {
  $('.wfa-clock-punchbutton').each(function() { 
    btnMap[$(this).data('punch-action')] = $(this);
  });
};

var source;
if(window.EventSource) {
  source = new window.EventSource('http://localhost:8080/commands', {withCredentials: true}) 
  source.addEventListener('punchEvent', punchAction);
}

var postStatus = function (status) {
  console.log(actions);
};

$(document).ready(function() {
  $('body').css('background-color', '#999');
  console.log('running btn scrape');
  buildMap();
  timer = setInterval(function() {update()}, 3000);
  // punchAction('ClockOutForDay');
});