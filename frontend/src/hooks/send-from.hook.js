export const useSendForm = () => {
  const addError = (element) => {
      element.parentElement.classList.add('inp-group--error')
  }
  const removeError = (element) => {
      element.parentElement.classList.remove('inp-group--error')
  }
  const checkPhone = (element) => {
      if (element.value.indexOf('_') !== -1 || element.value.length !== 18) {
          addError(element)
          return false
      }
      removeError(element)
      return true
  }

  const checkEmail = (element) => {
      const re = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
      if (!re.test(String(element.value).toLowerCase())) {
          addError(element)
          return false
      } else {
          removeError(element)
          return true
      }

  }
  const getUtms = (paths) => {
      let utm = {}
      if (paths.split('&').length > 0) {
          let params = paths.split('&')
          params.forEach((param) => {
              param = param.split('=')
              if (param[0] === "utm_medium" || param[0] === "utm_content" || param[0] === "utm_campaign" || param[0] === "utm_term" || param[0] === "utm_source") {
                  utm = { ...utm, [param[0]]: param[1] }
              }
          })
      }
      return utm
  }



  const showAlert = (data, celtype) => {
      if (window.Ya !== undefined) {
          let metriks = parseInt(Object.keys(window.Ya._metrika.counters)[0].split(':')[0])
          try {
              ym(metriks, 'reachGoal', celtype);
              ym(metriks, 'reachGoal', 'vse');
          } catch (e) { }
      }
      try {
          gtag('event', 'send', {
              'event_category': 'forms',
              'event_label': celtype,
              'value': 1
          });
          gtag('event', 'send', {
              'event_category': 'forms',
              'event_label': 'vse',
              'value': 1
          });
      } catch (e) { }
      try {
          fbq('trackCustom', 'vse');
      } catch (e) { }
      try {
          VK.Goal('lead')
      } catch (e) { }

      fetch("fd_log/lrcnt_react.php", {})
          .then(data => data.ok && data.json())
          .then(response => {
              data = { ...data, ...response }
              const requestOptions = {
                  method: 'POST',
                  headers: {
                      'Content-Type': 'application/json',
                      "Access-Control-Allow-Origin": "*",
                  },
                  body: JSON.stringify(data)
              };
              fetch("main/ajax.php", requestOptions)
                  .then(data => data.ok)
          });
  }

  const sendForm = (e, callback) => {
      e.preventDefault()
      let flag = true
      let form = e.target.closest('form')
      let celtype = e.currentTarget.getAttribute('celtype')
      let name = form.querySelector('input[name="name"]') !== null ? form.querySelector('input[name="name"]').value : ''
      let phone = form.querySelector('input[name="phone"]').value
      let url = window.location.toString().split("?")
      let email
      if (form.querySelector('input[name="email"]') !== null) {
          email = form.querySelector('input[name="email"]').value
      }
      let utm = null
      if (url.length > 1) {
          utm = getUtms(url[1])
      }

      if (!checkPhone(form.querySelector('input[name="phone"]'))) {
          flag = false
      }

      /*if (form.querySelector('input[name="name"]') !== null) {
          if (form.querySelector('input[name="name"]').value === "") {
              flag = false
              addError(form.querySelector('input[name="name"]'))
          } else {
              removeError(form.querySelector('input[name="name"]'))
          }
      }*/

      if (form.querySelector('input[name="email"]') !== null && form.querySelector('input[name="email"]').value !== "") {
          if (!checkEmail(form.querySelector('input[name="email"]'))) {
              flag = false
          }
      }
      let text = form.querySelector('input.text').value
      form.querySelectorAll('.dop-info').forEach((el) => {
          if (el.value.length !== 0 && el.getAttribute('data') !== null)
              text += ';' + el.getAttribute('data') + ':' + el.value
      })
      form.querySelectorAll('[name="dop-info"]').forEach((el) => {
          if (el.value.length !== 0)
              text += '; ' + el.value
      })
      form.querySelectorAll('.dop-select').forEach((el) => {
          if (el.value.length !== 0 && el.getAttribute('data')!== el.value)
              text += '; ' + el.getAttribute('data') + ":" + el.value
      })

      if (flag) {
          let senddata = { getCall: "Y", celtype: celtype, name: name, phone: phone, text: text, ...utm }
          if (email !== undefined) {
              senddata = { ...senddata, email: email }
          }
          const requestOptions = {
              method: 'POST',
              headers: {
                  'Content-Type': 'application/json',
                  "Access-Control-Allow-Origin": "*",
              },
              body: JSON.stringify(senddata)
          };
          document.querySelectorAll('.pu_inner').forEach((elem) => {
              elem.style.display = "none"
          })
          callback()
          console.log(senddata)
          /*fetch("fd_log/ajax.php", requestOptions)
              .then(data => data.ok)
              .then(response => {
                  showAlert(senddata, celtype)
              });*/
      }
  }
  return { sendForm }
}
