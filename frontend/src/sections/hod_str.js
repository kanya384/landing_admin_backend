import { useState } from "react";
import Slider from "react-slick/lib/slider";

export const HodStr = () => {
  const [nav1, setNav1] = useState();
  const [nav2, setNav2] = useState();

  let settings = {
    slidesToShow: 3,
    slidesToScroll: 1,
    dots: false,
    variableWidth: true,
    centerMode: true,
    arrows: true,
    infinite: true,
    centerPadding: 0,
    responsive: [
      {
        breakpoint: 1023,
        settings: {
          slidesToShow: 1,
          slidesToScroll: 1,
          centerMode: true,
          arrows: false,
          centerPadding: 10,
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
    focusOnSelect: true,
    infinite: true,
  }

  return (
    <div className="lvl14 parrallax-scene3-end">
      <div className="wrapper">
        <div className="construction-progress">
          <div className="title">Ход строительства</div>
          <div className="construction-progress__filters">
            <div className="filter-left">
              <button className="filter-link active">Фото</button>
              <button className="filter-link">Видео</button>
            </div>
            <div className="filter-right">
              <div className="filter-years">
                <button className="filter-link active">2021</button>
                <button className="filter-link">2022</button>
              </div>
              
              <div className="dropdown-filter js-dropdown-filter">
                <div className="dropdown-filter__label">
                  <div className="dropdown-filter__label-text">март</div>
                  <div className="dropdown-filter__arr">
                    <svg width="14" height="9" viewBox="0 0 14 9" xmlns="http://www.w3.org/2000/svg">
                      <path d="M13 1L7 7L1 1" strokeWidth="2"/>
                    </svg>
                  </div>
                </div>
                <div className="dropdown-filter__drop">
                  <ul className="dropdown-filter__list">
                    <li><span>март</span></li>
                    <li><span>апрель</span></li>
                    <li><span>май</span></li>
                    <li><span>июнь</span></li>
                  </ul>
                </div>
              </div>
            </div>
          </div>
          <div className="construction-progress__slider">
            <div className="slider-construction-progress js-slider-construction-progress-wrapper">
              <Slider asNavFor={nav2} ref={(slider1) => setNav1(slider1)} className="slider slider-construction-progress__slider-for js-slider-construction-progress__slider-for" {...settings}>
                <div className="slide"><img src="img/constr-prog1.png" alt=""/></div>
                <div className="slide"><img src="img/constr-prog1.png" alt=""/></div>
                <div className="slide"><img src="img/constr-prog1.png" alt=""/></div>
                <div className="slide"><img src="img/constr-prog1.png" alt=""/></div>

              </Slider>
              <div className="slider-construction-progress-nav">
                <div className="arr-left" onClick={()=> nav2.slickPrev()}>
                  <svg width="20" height="20" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
                    <path d="M1.5 10.5L0.792893 9.79289L0.0857865 10.5L0.792893 11.2071L1.5 10.5ZM8.20711 15.7929L2.20711 9.79289L0.792893 11.2071L6.79289 17.2071L8.20711 15.7929ZM2.20711 11.2071L8.20711 5.20711L6.79289 3.79289L0.792893 9.79289L2.20711 11.2071ZM1.5 11.5L18.5 11.5V9.5L1.5 9.5L1.5 11.5Z" />
                  </svg>
                </div>
                <div className="nav-build">
                  <Slider asNavFor={nav1} ref={(slider2) => setNav2(slider2)} className={"slider-construction-progress__slider-nav js-slider-construction-progress__slider-nav"} {...settings2}>

                    <div className="slider-nav-item"><span>01</span></div>
                    <div className="slider-nav-item"><span>02</span></div>
                    <div className="slider-nav-item"><span>03</span></div>
                    <div className="slider-nav-item"><span>04</span></div>
                  </Slider>
                </div>
                <div className="arr-right" onClick={()=> nav2.slickNext()}>
                  <svg width="20" height="20" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
                    <path d="M18.5 9.5L19.2071 10.2071L19.9142 9.5L19.2071 8.79289L18.5 9.5ZM11.7929 4.20711L17.7929 10.2071L19.2071 8.79289L13.2071 2.79289L11.7929 4.20711ZM17.7929 8.79289L11.7929 14.7929L13.2071 16.2071L19.2071 10.2071L17.7929 8.79289ZM18.5 8.5L1.5 8.5L1.5 10.5L18.5 10.5V8.5Z" />
                  </svg>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}