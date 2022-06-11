import React from "react"
import { AppBar } from "../components/appbar"
import { Footer } from "../components/footer"

export const PlanDetail = () => {
  return (
    <React.Fragment>
      <AppBar subfolder={true} />
      <div className="wrapper">
          <div className="page-title"><h1 className="h1">однокомнатная квартира №3</h1></div>
          <div className="content-wrapper">
            <main className="content">
              <div className="room-card-img-mobile">
                <span><img src="../../img/big-room.png" alt="" /></span>
              </div>
              <div className="options-room">
                <div className="option-room__item-wrap">
                  <div className="option-room__item">
                    <div className="label">Литер</div>
                    <div className="value">литер 3</div>
                  </div>
                </div>
                <div className="option-room__item-wrap">
                  <div className="option-room__item">
                    <div className="label">Подъезд</div>
                    <div className="value">1</div>
                  </div>
                </div>
                <div className="option-room__item-wrap">
                  <div className="option-room__item">
                    <div className="label">Общая площадь</div>
                    <div className="value">52.3 м²</div>
                  </div>
                </div>
                <div className="option-room__item-wrap">
                  <div className="option-room__item">
                    <div className="label">Сдача</div>
                    <div className="value">1 кв 2023</div>
                  </div>
                </div>
                <div className="option-room__item-wrap">
                  <div className="option-room__item">
                    <div className="label">Этаж</div>
                    <div className="value">1 из 9</div>
                  </div>
                </div>
                <div className="option-room__item-wrap">
                  <div className="option-room__item">
                    <div className="label">Жилая площадь</div>
                    <div className="value">29.3 м²</div>
                  </div>
                </div>
              </div>
              <div className="room-card">
                <div className="room-card__img">
                  <span><img src="../../img/big-room.png" alt="" /></span>
                </div>
                <div className="room-card__info">
                  <div className="room-card__title">Узнать стоимость квартиры на сегодня</div>
                  <div className="room-card__form form-ec__content">
                    <div className="form-ec__input-row">
                      <div className="inp-group">
                        <div className="inp-group-label">Ваше имя</div>
                        <input className="input" placeholder="Ваше имя" />
                      </div>
                      <div className="inp-group">
                        <div className="inp-group-label">Номер телефона</div>
                        <input className="input" placeholder="Номер телефона" />
                      </div>
                    </div>
                    <div className="form-ec__b-btn">
                      <button className="btn-submit"><span className="btn-submit-text">Узнать стоимость</span></button>
                    </div>
                  </div>
                </div>
              </div>
            </main>
          </div>
        </div>
        <Footer subfolder={true} />
    </React.Fragment>
  )
}