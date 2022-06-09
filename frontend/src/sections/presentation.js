import React, { useState } from "react"
import { Modal } from "../components/modals"

export const Presentation = () => {
  const [isOpen, setModalState] = useState(null)
  return ( <React.Fragment>
    <div class="lvl7">
      <div class="wrapper">
        <div class="form-ec">
          <div class="form-ec__title">Получите подробную презентацию жилого комплекса</div>
          <div class="form-ec__content">
            <div class="form-ec__input-row">
              <div class="inp-group">
                <div class="inp-group-label">Ваше имя</div>
                <input class="input" placeholder="Ваше имя" type="text" />
              </div>
              <div class="inp-group">
                <div class="inp-group-label">E-mail</div>
                <input class="input" placeholder="E-mail" type="text" />
              </div>
            </div>
            <div class="form-ec__b-btn"><button class="btn-submit" onClick={(e) => {e.preventDefault(); setModalState(true)}}><span class="btn-submit-text">Отправить</span></button></div>
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
      image ={"img/get-present-bg.svg"}
      imageMobile={"img/get-present-bg-mobile.svg"}
      opened={isOpen}
      close = {()=>{setModalState(null)}}
    />
    </React.Fragment>
  )
}