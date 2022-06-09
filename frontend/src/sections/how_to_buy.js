import React, { useState } from "react"
import { Modal } from "../components/modals"

export const HowToBuy = () => {
  const [modalType, setModalType] = useState(false)
  return (
    <React.Fragment>
        <div class="lvl9">
          <div class="wrapper">
            <div class="b-purch">
              <div class="b-purch-left">
                <div class="title">Способы покупки</div>
                <div class="b-purch-circ">
                  <div class="text">Более 50 банков-партнеров</div>
                </div>
                <div class="b-purch-circ">
                  <div class="text">Строим по эскроу-счетам</div>
                </div>
              </div>
              <div class="b-purch-right">
                <div class="b-purch-list">
                  <div class="b-purch-item">
                    <span class="text">Выберите подходящий вариант оплаты, а мы поможем оформить сделку</span>
                  </div>
                  <div class="b-purch-item b-purch-item--mobile">
                    <div class="b-purch-item__row">
                      <div class="b-purch-item__col-50 gray-text">
                        Более 50 банков-партнеров
                      </div>
                      <div class="b-purch-item__col-50 gray-text">
                        Строим по эскроу-счетам
                      </div>
                    </div>
                  </div>
                  <a href="" class="b-purch-item" onClick={(e)=>{e.preventDefault(); setModalType(0)}}>
                    <span class="title">Ипотека от 6%</span>
                    <span class="text">Мы помогаем подать заявку сразу в несколько ведущих банков</span>
                  </a>
                  <a href="" class="b-purch-item" onClick={(e)=>{e.preventDefault(); setModalType(1)}}>
                    <span class="title">Рассрочка 0%</span>
                    <span class="text">Оформляем рассрочку на 6 месяцев. Первоначальный взнос - от 70%</span>
                  </a>
                  <a href="" class="b-purch-item" onClick={(e)=>{e.preventDefault(); setModalType(2)}}>
                    <span class="title">Материнский капитал</span>
                    <span class="text">В качестве первого взноса</span>
                  </a>
                  <a href="" class="b-purch-item" onClick={(e)=>{e.preventDefault(); setModalType(3)}}>
                    <span class="title">Военная ипотека</span>
                    <span class="text">Специальные условия для военных</span>
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