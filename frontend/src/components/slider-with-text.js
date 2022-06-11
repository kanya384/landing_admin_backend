import { useEffect, useState } from "react";
import Slider from "react-slick/lib/slider";

export const SliderWithText = ({slides}) => {
  const [nav1, setNav1] = useState();
  const [nav2, setNav2] = useState();


  let settings = {
    slidesToShow: 1,
    slidesToScroll: 1,
    fade: true,
    dots: false,
    arrows: false,
    variableWidth: true,
    infinity: true,
  }

  let settings2 = {
    slidesToShow: 3,
    slidesToScroll: 1,
    dots: false,
    arrows: false,
    vertical: false,
    infinity: true,
  }


  return (
    <div className="about-info-item">
                  <div className="img">
                  <div className="slider-vert js-slider-vert-wrapper">
                    <Slider asNavFor={nav2} ref={(slider1) => setNav1(slider1)} className="slider slider-vert__slider-for js-slider-vert__slider-for" {...settings}>
                      {slides.slides.map((slide)=>{
                        return <div className="slide"><img src={slide} alt="" /></div>
                      })}
                      {slides.slides.map((slide)=>{
                        return <div className="slide"><img src={slide} alt="" /></div>
                      })}
                    </ Slider>
                    <div className="slider-vert-nav">
                      <div className="arr-left" onClick={()=>nav1.slickPrev()}>
                        <svg width="20" height="20" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
                          <path d="M1.5 10.5L0.792893 9.79289L0.0857865 10.5L0.792893 11.2071L1.5 10.5ZM8.20711 15.7929L2.20711 9.79289L0.792893 11.2071L6.79289 17.2071L8.20711 15.7929ZM2.20711 11.2071L8.20711 5.20711L6.79289 3.79289L0.792893 9.79289L2.20711 11.2071ZM1.5 11.5L18.5 11.5V9.5L1.5 9.5L1.5 11.5Z" />
                        </svg>
                      </div>
                      <div className="nav-build">
                          <Slider asNavFor={nav1} ref={(slider2) => setNav2(slider2)} className="slider-vert__slider-nav js-slider-vert__slider-nav" {...settings2}>
                            {Array.from(Array(slides.slides.length).keys()).map((index)=>{
                              return <div className="slider-nav-item"><span>{index+1<10?"0"+(index+1):index+1}</span></div>
                            })}
                            {Array.from(Array(slides.slides.length).keys()).map((index)=>{
                              return <div className="slider-nav-item"><span>{index+1<10?"0"+(index+1):index+1}</span></div>
                            })}
                          </Slider>
                      </div>
                      <div className="arr-right" onClick={()=>nav1.slickNext()}>
                        <svg width="20" height="20" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
                          <path d="M18.5 9.5L19.2071 10.2071L19.9142 9.5L19.2071 8.79289L18.5 9.5ZM11.7929 4.20711L17.7929 10.2071L19.2071 8.79289L13.2071 2.79289L11.7929 4.20711ZM17.7929 8.79289L11.7929 14.7929L13.2071 16.2071L19.2071 10.2071L17.7929 8.79289ZM18.5 8.5L1.5 8.5L1.5 10.5L18.5 10.5V8.5Z" />
                        </svg>
                      </div>
                    </div>
                  </div>
                  </div>
                  <div className="b-text">
                    <div className="title">{slides.title}</div>
                    <div className="text">{slides.text}</div>
                  </div>
                </div>
  )
}