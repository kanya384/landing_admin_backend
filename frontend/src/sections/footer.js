export const Footer = () => {
  return (<footer class="footer">
      <div class="wrapper">
          <div class="footer-row">
              <div class="footer-left">
                  <div class="footer-map"><img src="img/contact-map.svg" alt="" /></div>
                  <div class="footer-logo"><a href=""><img src="img/foot-logo.svg" alt="" /></a></div>
              </div>
              <div class="footer-right">
                  <div class="title">Или свяжитесь с нами прямо сейчас:  </div>
                  <div class="footer-contact-info">
                      <div class="b-col b-col--addr text-uppercase">
                          <strong>г. Анапа, <br />ул. Крылова, 13К1</strong>
                      </div>
                      <div class="b-col b-col--phone">
                          <strong>+7 (987) 654-32-10</strong>
                      </div>
                      <div class="b-col b-col--time">
                          пн-пт с 9:00 до 18:00, <br />
                          сб-вс с 10:00 до 16:00
                      </div>
                  </div>
                  <div class="footer-text">
                      Отправляя любую форму на сайте, вы соглашаетесь с политикой конфиденциальности данного сайта, а также на получение рассылки на электронную почту, указанную в заявку
                  </div>
              </div>
          </div>
          <div class="footer-logo--mobile"><a href=""><img src="img/foot-logo.svg" alt="" /></a></div>
      </div>
  </footer>
  )
}