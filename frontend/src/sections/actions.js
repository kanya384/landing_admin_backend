import React, { useState } from "react"
import Slider from "react-slick/lib/slider";
import { Modal } from "../components/modals"

export const Actions = () => {
    const [nav1, setNav1] = useState();
    const [nav2, setNav2] = useState();
    const [info, setInfo] = useState({
        date: "",
        title: "",
        description:"",
        img: "",
    })

    const [isOpen, setModalState] = useState(null)

    let settings = {
        slidesToShow: 3,
        slidesToScroll: 1,
        dots: false,
        // variableWidth: true,
        // centerMode: true,
        arrows: false,
        responsive: [
            {
                breakpoint: 769,
                settings: {
                    slidesToShow: 2,
                    slidesToScroll: 1,
                }
            },
            {
                breakpoint: 560,
                settings: {
                    rows: 3,
                    adaptiveHeight: true,
                    slidesToShow: 1,
                    slidesToScroll: 1,
                }
            },
        ]
    }
    
    let settings2 = {
        slidesToShow: 3,
        slidesToScroll: 1,
        dots: false,
        arrows: false,
        vertical: false,
        infinity: true,
    }

    let infoTest = {
        date: "До 22 августа",
        title: "Акция «Весна в Красной Поляне»",
        description: "<p>Отделка первого этажа выполнена клинкерной плиткой, со 2 по 9 этаж вентилируемый фасад с использованием крупно-форматных алюминиевых кассет и декоративными вставками. Для сохранения единого фасадного стиля в процессе эксплуатации используются металлические корзины для кондиционеров, а также остекление лоджий профилем в стиль фасада.</p><p>Много естественного света в подъездах обеспечат большие стеклянные двери премиум-сегмента.</p>",
    } 

    return <React.Fragment>
        <div className="lvl-stock">
            <div className="wrapper">
                <h2 className="lvl-title">акции</h2>
                <div className="slider-stock js-slider-stock-wrapper">
                <Slider asNavFor={nav2} ref={(slider1) => setNav1(slider1)} className="slider slider-stock__slider-for js-slider-stock__slider-for" {...settings}>
                    <div className="slide">
                        <div className="stock-item" onClick={(e) => {e.preventDefault();  setInfo(infoTest); setModalState(true)}}>
                            <div className="stock-item__img"><a href="#"><img src="img/akcii1.png" alt=""/></a></div>
                            <div className="stock-item__right">
                            <div className="stock-item__title"><a href="#">Акция «Весна в Красной Поляне»</a></div>
                            <div className="stock-item__data">До 22 августа</div>
                            </div>
                        </div>
                    </div>
                    <div className="slide">
                        <div className="stock-item">
                            <div className="stock-item__img"><a href="#"><img src="img/akcii1.png" alt=""/></a></div>
                            <div className="stock-item__right">
                                <div className="stock-item__title"><a href="#">Рассрочка на квартиры в Сочи и в Челябинске</a></div>
                                <div className="stock-item__data">До 22 августа</div>
                            </div>
                        </div>
                    </div>
                    <div className="slide">
                        <div className="stock-item">
                            <div className="stock-item__img"><a href="#"><img src="img/akcii3.png" alt=""/></a></div>
                            <div className="stock-item__right">
                                <div className="stock-item__title"><a href="#">Акция «Весна в Красной Поляне»</a></div>
                                <div className="stock-item__data">До 22 августа</div>
                            </div>
                        </div>
                    </div>
                    <div className="slide">
                        <div className="stock-item">
                            <div className="stock-item__img"><a href="#"><img src="img/akcii2.png" alt=""/></a></div>
                            <div className="stock-item__right">
                                <div className="stock-item__title"><a href="#">Рассрочка на квартиры в Сочи и в Челябинске</a></div>
                                <div className="stock-item__data">До 22 августа</div>
                            </div>
                        </div>
                    </div>
                    <div className="slide">
                        <div className="stock-item">
                            <div className="stock-item__img"><a href="#"><img src="img/akcii3.png" alt=""/></a></div>
                            <div className="stock-item__right">
                                <div className="stock-item__title"><a href="#">Акция «Весна в Красной Поляне»</a></div>
                                <div className="stock-item__data">До 22 августа</div>
                            </div>
                        </div>
                    </div>
                    <div className="slide">
                        <div className="stock-item">
                            <div className="stock-item__img"><a href="#"><img src="img/akcii2.png" alt=""/></a></div>
                            <div className="stock-item__right">
                                <div className="stock-item__title"><a href="#">Рассрочка на квартиры в Сочи и в Челябинске</a></div>
                                <div className="stock-item__data">До 22 августа</div>
                            </div>
                        </div>
                    </div>
                    <div className="slide">
                        <div className="stock-item">
                            <div className="stock-item__img"><a href="#"><img src="img/akcii3.png" alt=""/></a></div>
                            <div className="stock-item__right">
                                <div className="stock-item__title"><a href="#">Акция «Весна в Красной Поляне»</a></div>
                                <div className="stock-item__data">До 22 августа</div>
                            </div>
                        </div>
                    </div>
                    <div className="slide">
                        <div className="stock-item">
                            <div className="stock-item__img"><a href="#"><img src="img/akcii1.png" alt=""/></a></div>
                            <div className="stock-item__right">
                                <div className="stock-item__title"><a href="#">Акция «Весна в Красной Поляне»</a></div>
                                <div className="stock-item__data">До 22 августа</div>
                            </div>
                        </div>
                    </div>
                </Slider>
                <div className="slider-stock-nav">
                    <div className="arr-left" onClick={()=>nav1.slickPrev()}>
                        <svg width="20" height="20" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
                            <path d="M1.5 10.5L0.792893 9.79289L0.0857865 10.5L0.792893 11.2071L1.5 10.5ZM8.20711 15.7929L2.20711 9.79289L0.792893 11.2071L6.79289 17.2071L8.20711 15.7929ZM2.20711 11.2071L8.20711 5.20711L6.79289 3.79289L0.792893 9.79289L2.20711 11.2071ZM1.5 11.5L18.5 11.5V9.5L1.5 9.5L1.5 11.5Z" />
                        </svg>
                    </div>
                    <div className="arr-right" onClick={()=>nav1.slickNext()}>
                        <svg width="20" height="20" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
                            <path d="M18.5 9.5L19.2071 10.2071L19.9142 9.5L19.2071 8.79289L18.5 9.5ZM11.7929 4.20711L17.7929 10.2071L19.2071 8.79289L13.2071 2.79289L11.7929 4.20711ZM17.7929 8.79289L11.7929 14.7929L13.2071 16.2071L19.2071 10.2071L17.7929 8.79289ZM18.5 8.5L1.5 8.5L1.5 10.5L18.5 10.5V8.5Z" />
                        </svg>
                    </div>
                </div>
                </div>
            </div>
        </div>
        <Modal 
            title={info.title}
            content={info.description}
            date={info.date}
            classes={"modal-stock"}
            position={window.pageYOffset}
            fields={[]}
            celtype={""}
            image ={"img/action_big.png"}
            opened={isOpen}
            close = {()=>{setModalState(null)}}
        />
    </React.Fragment>
}

export default Actions