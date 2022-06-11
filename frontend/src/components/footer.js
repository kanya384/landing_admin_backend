export const Footer = ({subfolder}) => {
  return (<footer className="footer">
      <div className="wrapper">
          <div className="footer-row">
              <div className="footer-left">
                  <div className="footer-map"><img src={subfolder?"../../img/contact-map.svg":"img/contact-map.svg"} alt="" /></div>
                  <div className="footer-logo"><a href=""><img src={subfolder?"../../img/foot-logo.svg":"img/foot-logo.svg"} alt="" /></a></div>
              </div>
              <div className="footer-right">
                  <div className="title">Или свяжитесь с нами прямо сейчас:  </div>
                  <div className="footer-contact-info">
                      <div className="b-col b-col--addr text-uppercase">
                          <strong>г. Анапа, <br />ул. Крылова, 13К1</strong>
                      </div>
                      <div className="b-col b-col--phone">
                          <strong>+7 (987) 654-32-10</strong>
                      </div>
                      <div className="b-col b-col--time">
                          пн-пт с 9:00 до 18:00, <br />
                          сб-вс с 10:00 до 16:00
                      </div>
                  </div>
                  <div className="footer-text">
                      Отправляя любую форму на сайте, вы соглашаетесь с политикой конфиденциальности данного сайта, а также на получение рассылки на электронную почту, указанную в заявку
                  </div>
              </div>
          </div>
          <div className="footer-logo--mobile"><a href=""><img src={subfolder?"../../img/foot-logo.svg":"img/foot-logo.svg"} alt="" /></a></div>
      </div>
  </footer>
  )
}