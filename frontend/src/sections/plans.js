import React, { useEffect, useState } from "react"
import { Modal } from "../components/modals"
import { Link } from "react-router-dom";

export const Plans = () => {
  const [isOpen, setModalState] = useState(false)

  return (
    <React.Fragment>
      <div className="lvl8">
        <div className="wrapper">
          <div className="h-title">Подбор квартиры на 3D-плане</div>
          <div className="text">Выберите дом и этаж, чтобы посмотреть планировки и узнать цены. Передвигайтесь влево-вправо, ввех и вниз</div>
          <div className="b-3d" onClick={() => setModalState(true)}>
            <img src="img/home3d.svg" alt="" />
          </div>
          <div className="b-link-row">
            <Link to="/plans" className="lnk-params js-open-modal-flat">
              <span className="lnk-params__ico">
              <svg className="ico" width="66" height="30" viewBox="0 0 66 30"
                    xmlns="http://www.w3.org/2000/svg">
                <path d="M64.5 14.9983L65.2033 15.7092L65.9219 14.9982L65.2032 14.2874L64.5 14.9983ZM51.0509 29.7109L65.2033 15.7092L63.7967 14.2875L49.6443 28.2891L51.0509 29.7109ZM65.2032 14.2874L51.0508 0.289034L49.6444 1.71097L63.7968 15.7093L65.2032 14.2874ZM64.5 13.9983H0.5V15.9983H64.5V13.9983Z" />
              </svg>
              </span>
              Подбор по параметрам
            </Link>
          </div>
        </div>
      </div>
      <Modal 
        title={"Планировки литера 2"}
        position={window.pageYOffset}
        classes={"modal-flat"}
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