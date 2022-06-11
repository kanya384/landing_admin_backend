import React, { useState } from "react"
import { Modal } from "../components/modals"

export const HowToBuy = () => {
  const [modalType, setModalType] = useState(false)
  return (
    <React.Fragment>
        <div className="lvl9">
          <div className="wrapper">
            <div className="b-purch">
              <div className="b-purch-left">
                <div className="title">Способы покупки</div>
                <div className="b-purch-circ">
                  <div className="text">Более 50 банков-партнеров</div>
                </div>
                <div className="b-purch-circ">
                  <div className="text">Строим по эскроу-счетам</div>
                </div>
              </div>
              <div className="b-purch-right">
                <div className="b-purch-list">
                  <div className="b-purch-item">
                    <span className="text">Выберите подходящий вариант оплаты, а мы поможем оформить сделку</span>
                  </div>
                  <div className="b-purch-item b-purch-item--mobile">
                    <div className="b-purch-item__row">
                      <div className="b-purch-item__col-50 gray-text">
                        Более 50 банков-партнеров
                      </div>
                      <div className="b-purch-item__col-50 gray-text">
                        Строим по эскроу-счетам
                      </div>
                    </div>
                  </div>
                  <a href="" className="b-purch-item" onClick={(e)=>{e.preventDefault(); setModalType(0)}}>
                    <span className="title">Ипотека от 6%</span>
                    <span className="text">Мы помогаем подать заявку сразу в несколько ведущих банков</span>
                  </a>
                  <a href="" className="b-purch-item" onClick={(e)=>{e.preventDefault(); setModalType(1)}}>
                    <span className="title">Рассрочка 0%</span>
                    <span className="text">Оформляем рассрочку на 6 месяцев. Первоначальный взнос - от 70%</span>
                  </a>
                  <a href="" className="b-purch-item" onClick={(e)=>{e.preventDefault(); setModalType(2)}}>
                    <span className="title">Материнский капитал</span>
                    <span className="text">В качестве первого взноса</span>
                  </a>
                  <a href="" className="b-purch-item" onClick={(e)=>{e.preventDefault(); setModalType(3)}}>
                    <span className="title">Военная ипотека</span>
                    <span className="text">Специальные условия для военных</span>
                  </a>
                </div>
              </div>
            </div>
          </div>
        </div>
        <Modal
            title={"рассчитать ипотеку"}
            subtitle={"Уникальные ипотечные программы и специальные условия от ведущих банков.<br /> Всего один шаг, и квартира- Ваша."}
            classes={"modal-mortgage-calculation"}
            position={window.pageYOffset}
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
            btnTitle={"Получить расчет"}
            celtype={"getIpot"}
            opened={modalType===0}
            close = {()=>{setModalType(null)}}
          />
        <Modal
            title={"рассчитать рассрочку"}
            //subtitle={"Уникальные ипотечные программы и специальные условия от ведущих банков.<br /> Всего один шаг, и квартира- Ваша."}
            classes={"modal-mortgage-calculation"}
            position={window.pageYOffset}
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
            btnTitle={"Получить расчет"}
            celtype={"getIpot"}
            opened={modalType===1}
            close = {()=>{setModalType(null)}}
          />
        <Modal
            title={"материнский капитал"}
            //subtitle={"Уникальные ипотечные программы и специальные условия от ведущих банков.<br /> Всего один шаг, и квартира- Ваша."}
            classes={"modal-mortgage-calculation"}
            position={window.pageYOffset}
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
            btnTitle={"Получить консультацию"}
            celtype={"getMatKap"}
            opened={modalType===2}
            close = {()=>{setModalType(null)}}
          />

        <Modal
            title={"военная ипотека"}
            //subtitle={"Уникальные ипотечные программы и специальные условия от ведущих банков.<br /> Всего один шаг, и квартира- Ваша."}
            classes={"modal-mortgage-calculation"}
            position={window.pageYOffset}
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
            btnTitle={"Получить консультацию"}
            celtype={"getVoenIpot"}
            opened={modalType===3}
            close = {()=>{setModalType(null)}}
          />

    </React.Fragment>
  )
}