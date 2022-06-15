import React, { useState } from "react"
import { Modal } from "../components/modals"

export const Presentation = () => {
  const [isOpen, setModalState] = useState(null)
  return ( <React.Fragment>
    <div className="lvl7">
      <div className="wrapper">
        <div className="form-ec">
          <div className="form-ec__title">Получите подробную презентацию жилого комплекса</div>
          <div className="form-ec__content">
            <div className="form-ec__input-row">
              <div className="inp-group">
                <div className="inp-group-label">Ваше имя</div>
                <input className="input" placeholder="Ваше имя" type="text" />
              </div>
              <div className="inp-group">
                <div className="inp-group-label">E-mail</div>
                <input className="input" placeholder="E-mail" type="text" />
              </div>
            </div>
            <div className="form-ec__b-btn"><button className="btn-submit" onClick={(e) => {e.preventDefault(); setModalState(true)}}><span className="btn-submit-text">Отправить</span></button></div>
          </div>
        </div>
      </div>
    </div>
    <Modal 
      title={"получить презентацию"}
      position={window.pageYOffset}
      classes={"modal-get-present"}
      fields={[
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
      btnTitle={"Отправить"}
      celtype={"getPresentation"}
      image ={"img/get-present-bg.jpg"}
      imageMobile={"img/get-present-bg-mobile.jpg"}
      opened={isOpen}
      close = {()=>{setModalState(null)}}
    />
    </React.Fragment>
  )
}