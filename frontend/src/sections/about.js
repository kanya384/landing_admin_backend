import React, { useState } from "react"
import { Modal } from "../components/modals"

export const  About = () => {
  const [isOpen, setModalState] = useState(null)
  return ( <React.Fragment>
          <div className="lvl6">
            <div className="wrapper">
              <div className="about-title">О проекте</div>
              <div className="about-text">
                3 девятиэтажных дома на Высоком берегу по своему архитектурному облику напоминают волны, среди которых расположится благоустройство и озеленение жилого комплекса, в традиционно высоком исполнении
              </div>
              <div className="about-info-list">
                <div className="about-info-item">
                  <div className="img"><img src="img/img6.svg" alt="" /></div>
                  <div className="b-text">
                    <div className="title">Безопасный и умный дом</div>
                    <div className="text">Один из главных приоритетов жизни в жилом комплексе AVAnta - комфорт и безопасность жителей</div>
                    <div className="podr"><a href="" onClick={(e) => {e.preventDefault(); setModalState(true)}}>Подробнее</a></div>
                  </div>
                </div>
                <div className="about-info-item">
                  <div className="img"><img src="img/img7.svg" alt="" /></div>
                  <div className="b-text">
                    <div className="title">Двор как место притяжения</div>
                    <div className="text">Дворы дружбы и добрых соседей — именно такой принцип заложен в основу благоустройства территории</div>
                    <div className="podr"><a href="" onClick={(e) => {e.preventDefault(); setModalState(true)}}>Подробнее</a></div>
                  </div>
                </div>
                <div className="about-info-item">
                  <div className="img"><img src="img/img8.svg" alt="" /></div>
                  <div className="b-text">
                    <div className="title">Фасады AVAnta</div>
                    <div className="text">Вентилируеймый фасад жилого комплекса подчеркивает статус домов класса комфорт+</div>
                    <div className="podr"><a href="" onClick={(e) => {e.preventDefault(); setModalState(true)}}>Подробнее</a></div>
                  </div>
                </div>
                <div className="about-info-item">
                  <div className="img"><img src="img/img9.svg" alt="" /></div>
                  <div className="b-text">
                    <div className="title">Дом изнутри</div>
                    <div className="text">Отделка мест общественного пользования дополняет архитектуру зданий и выполнена в теплом морском стиле</div>
                    <div className="podr"><a href="" onClick={(e) => {e.preventDefault(); setModalState(true)}}>Подробнее</a></div>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <Modal 
            title={"Безопасный и умный дом"}
            content={"<p>Отделка первого этажа выполнена клинкерной плиткой, со 2 по 9 этаж вентилируемый фасад с использованием крупно-форматных алюминиевых кассет и декоративными вставками. Для сохранения единого фасадного стиля в процессе эксплуатации используются металлические корзины для кондиционеров, а также остекление лоджий профилем в стиль фасада.</p><p>Много естественного света в подъездах обеспечат большие стеклянные двери премиум-сегмента.</p>"}
            position={window.pageYOffset}
            fields={[]}
            celtype={""}
            image ={"img/news-img.svg"}
            imageMobile={"img/news-mobile.svg"}
            opened={isOpen}
            close = {()=>{setModalState(null)}}
        />
    </React.Fragment>
  )
}