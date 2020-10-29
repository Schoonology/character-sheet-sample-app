window.addEventListener('DOMContentLoaded', (event) => {
  var elDialog = document.querySelector('#result')
  var elResult = document.querySelector('#result .value')

  document
    .querySelector('#result button')
    .addEventListener('click', hideResult)

  wireUpButton('btnStrength', 'strength')
  wireUpButton('btnDexterity', 'dexterity')
  wireUpButton('btnConstitution', 'constitution')
  wireUpButton('btnIntelligence', 'intelligence')
  wireUpButton('btnWisdom', 'wisdom')
  wireUpButton('btnCharisma', 'charisma')

  function wireUpButton(id, input) {
    var elStat = document.querySelector(`input[name=${input}]`)

    document
      .querySelector(`#${id}`)
      .addEventListener('click', (event) => {
        var stat = Number(elStat.value)

        fetch('/roll/d20', {})
          .then(response => response.json())
          .then(data => showResult(data.result, modifierFor(stat)))
      })
  }

  function modifierFor(stat) {
    switch (stat) {
      case 6:
      case 7:
        return -2;
      case 8:
      case 9:
        return -1;
      case 10:
      case 11:
        return 0;
      case 12:
      case 13:
        return 1;
      case 14:
      case 15:
        return 2;
      case 16:
      case 17:
        return 3;
      case 18:
      case 19:
        return 4;
      case 20:
        return 5;
    }
  }

  function showResult(val, modifier) {
    elResult.innerHTML = `${val}${!modifier ? '' : ' + ' + modifier + ' = ' + (val + modifier)}`

    elDialog.classList.add('visible')
  }

  function hideResult() {
    elDialog.classList.remove('visible')
  }
})
