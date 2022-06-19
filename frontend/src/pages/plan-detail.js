import React, { useContext, useEffect, useState } from "react"
import { useParams } from "react-router-dom"
import { AppBar } from "../components/appbar"
import { Footer } from "../components/footer"
import { ContentContext } from "../context/contentContext"

export const PlanDetail = () => {
  const [flat, setFlat] = useState()
  const content = useContext(ContentContext)
  
  let { id } = useParams();
  useEffect(()=>{
    
    content.content?.Plans.forEach((flat)=>{
      if (flat.ID === id) {
        setFlat(flat)
        window.scrollTo(0,0)
        console.log(flat)
        return
      }
    })
  },[id, content.content])
  return (
    <React.Fragment>
      <AppBar subfolder={true} />
      <div className="wrapper">
          <div className="page-title"><h1 className="h1">{flat?.rooms === 1?"однокомнатная":flat?.rooms === 2?"двухкомнатная":"трехкомнатная"} квартира №{flat?.number}</h1></div>
          <div className="content-wrapper">
            <main className="content">
              <div className="room-card-img-mobile">
                <span><img src={process.env.REACT_APP_BACKEND+"store/"+flat?.image} alt="" /></span>
              </div>
              <div className="options-room">
                <div className="option-room__item-wrap">
                  <div className="option-room__item">
                    <div className="label">Литер</div>
                    <div className="value">литер {flat?.liter.replace('ЖК ВЫСОКИЙ БЕРЕГ Литер ',"")}</div>
                  </div>
                </div>
                <div className="option-room__item-wrap">
                  <div className="option-room__item">
                    <div className="label">Подъезд</div>
                    <div className="value">{flat?.entrance}</div>
                  </div>
                </div>
                <div className="option-room__item-wrap">
                  <div className="option-room__item">
                    <div className="label">Общая площадь</div>
                    <div className="value">{flat?.area} м²</div>
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
                    <div className="value">{flat?.floor} из 9</div>
                  </div>
                </div>
                <div className="option-room__item-wrap">
                  <div className="option-room__item">
                    <div className="label">Жилая площадь</div>
                    <div className="value">{flat?.living_area} м²</div>
                  </div>
                </div>
              </div>
              <div className="room-card">
                <div className="room-card__img">
                  <span><img src={process.env.REACT_APP_BACKEND+"store/"+flat?.image} alt="" /></span>
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