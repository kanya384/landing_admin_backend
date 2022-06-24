import { EditableText } from "./editable-text"
import { YMaps, Map, ZoomControl, Placemark } from 'react-yandex-maps';

export const Footer = ({subfolder}) => {
  return (<footer className="footer">
      <div className="wrapper">
          <div className="footer-row">
              <div className="footer-left">
                  <div className="footer-map" style={{width:"100%", height: 368, display:"block", position: "relative", overflow: "hidden" }}>
                        <YMaps>
                                <Map defaultState={{ center: [44.876247, 37.322943], zoom: 15, controls: [] }} style={{ width: "100%", height: "100%", position: "absolute" }}  >
                                    <ZoomControl options={{ float: 'left' }} />
                                    <Placemark geometry={[44.876247, 37.322943]}
                                                options={{
                                                    iconLayout: 'default#image',
                                                    hideIconOnBalloonOpen: false,
                                                    iconImageSize: [50, 50],
                                                    iconImageOffset: [-50, -72],
                                                    cursor: 'default',
                                                    iconShadow: true,
                                                    balloonclose: true,
                                                    iconImageHref: "img/icon-map-ava.svg",
                                                }}
                                                
                                            />
                                    
                                </Map>
                        </YMaps>
                  </div>
                  <div className="footer-logo"><a href=""><img src={subfolder?"../../img/foot-logo.svg":"img/foot-logo.svg"} alt="" /></a></div>
              </div>
              <div className="footer-right">
                  <div className="title"><EditableText id={"62aefdb322de07ce5e317519"} defaultText={"Или свяжитесь с нами прямо сейчас:"}/></div>
                  <div className="footer-contact-info">
                      <div className="b-col b-col--addr text-uppercase">
                          <strong><EditableText id={"62aefdb322de07ce5e31751a"} defaultText={"г. Анапа, <br />ул. Крылова, 13К1"}/></strong>
                      </div>
                      <div className="b-col b-col--phone">
                          <strong><EditableText id={"62aefdb322de07ce5e31751b"} defaultText={"+7 (987) 654-32-10"}/></strong>
                      </div>
                      <div className="b-col b-col--time">
                        <EditableText id={"62aefdb322de07ce5e31751c"} defaultText={"пн-пт с 9:00 до 18:00, <br /> сб-вс с 10:00 до 16:00"}/>
                      </div>
                  </div>
                  <div className="footer-text">
                      <EditableText id={"62aefdb322de07ce5e31751d"} defaultText={"Отправляя любую форму на сайте, вы соглашаетесь с политикой конфиденциальности данного сайта, а также на получение рассылки на электронную почту, указанную в заявку"}/>
                  </div>
              </div>
          </div>
          <div className="footer-logo--mobile"><a href=""><img src={subfolder?"../../img/foot-logo.svg":"img/foot-logo.svg"} alt="" /></a></div>
      </div>
  </footer>
  )
}

export default Footer