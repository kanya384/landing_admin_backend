import React, { useContext, useEffect, useState } from "react"
import { useParams } from "react-router-dom"
import { AppBar } from "../components/appbar"
import { Footer } from "../components/footer"
import { ContentContext } from "../context/contentContext"
import { Modal } from "../components/modals";
import { Form } from "../components/form";


export const PlanDetail = () => {
  const [isOpen, setIsOpen] = useState(false)
  const [flat, setFlat] = useState()
  const content = useContext(ContentContext)
  const [showPrices, setShowPrices] = useState(false)
  
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

  useEffect(()=>{
    if ( content.content){
      content.content.Setting.map((setting)=>{
        if (setting.name == "showPrices") {
          setShowPrices(setting.value == 1)
        }
      })
    }
  }, [content.content])
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
                  {showPrices?<div className="room-card__info-header">
                    <div className="room-card__info-price">54 000 000 р.</div>
                    <div className="room-card__info-text">
                      Остались вопросы? <br/>
                      Наши менеджеры помогут Вам!
                    </div>
                  </div>:<div className="room-card__title">Узнать стоимость квартиры на сегодня</div>}
                  <Form fields={[
                      {
                        type:"text",
                        name: "name",
                        placeholder: "Имя",
                        required: false,
                      },
                      {
                          type:"text",
                          name: "phone",
                          placeholder: "Телефон",
                          required: true,
                      },
                    ]}
                    btnTitle={`Узнать стоимость`}
                    description={`Узнать стоимость квартиры №${flat?.number}; Литер: ${flat?.liter}; Подъезд: ${flat?.entrance}; Этаж:${flat?.floor}; Кол-во комнат: ${flat?.rooms}; Общая площадь: ${flat?.area}`}
                    celtype={"getFlat"}
                    close={()=>{}}
                    callback={()=>{setIsOpen(true)}}
                />
                </div>
              </div>
            </main>
          </div>
        </div>
        <Footer subfolder={true} />
        <Modal
          success={true}
          position={window.pageYOffset}
          opened={isOpen}
          close = {()=>{setIsOpen(null)}}
        />
    </React.Fragment>
  )
}

export default PlanDetail
