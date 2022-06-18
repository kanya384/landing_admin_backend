import { useContext, useEffect, useState } from "react";
import Slider from "react-slick/lib/slider";
import { ContentContext } from "../context/contentContext";

export const Posters = () => {
  const [nav1, setNav1] = useState();
  const [nav2, setNav2] = useState();
  const [posters, setPosters] = useState([])
  const [index, setIndex] = useState(0)

  const content = useContext(ContentContext)
  
  const settings = {
    slidesToShow: 1,
    slidesToScroll: 1,
    fade: true,
    dots: false,
    arrows: true,
    arrows: false,
    beforeChange: (current, next) => setIndex(next)
  }

  const settings2 = {
    slidesToShow: 3,
    slidesToScroll: 1,
    dots: false,
    arrows: false,
    vertical: true,
    focusOnSelect: true
  }

  useEffect(()=>{
    if (posters.length === 0 && content.content !== undefined) {
      setPosters(content.content.Posters)
      setIndex(0)
      setTimeout(()=>nav1.slickNext(),200)
    }
  },[content])
 

  return <div className="lvl1">
      <div className="wrapper">
        <div className="row-lvl">
          <div className="slide-text">
            <div className="title">
              {posters[index] ? posters[index].title:"Жизнь на высоком берегу"}
            </div>
            <div className="text">
              Квартиры в Анапе от 5 млн руб.
            </div>
          </div>
          <div className="slider-lvl1">
            <div className="slider-horiz js-slider-hor-wrapper">
              <Slider asNavFor={nav2} ref={(slider1) => setNav1(slider1)} className="slider slider-lvl1__slider-for js-slider-hor__slider-for" {...settings}>
                {posters.length>0 ?posters.map((poster)=>{
                  return <div className="slide"><img src={process.env.REACT_APP_BACKEND+"store/"+ poster.photo} alt="" /></div>
                }):""}
              </Slider>
              <div className="slider-horiz-nav">
                <div className="arr-up" onClick={()=>nav1.slickPrev()}>
                  <svg width="14" height="19" viewBox="0 0 14 19" xmlns="http://www.w3.org/2000/svg">
                    <path d="M7 2L7.70711 1.29289L7 0.585786L6.29289 1.29289L7 2ZM1.70711 8.70711L7.70711 2.70711L6.29289 1.29289L0.292893 7.29289L1.70711 8.70711ZM6.29289 2.70711L12.2929 8.70711L13.7071 7.29289L7.70711 1.29289L6.29289 2.70711ZM6 2V19H8V2H6Z" />
                  </svg>
                </div>
                <div className="nav-build">
                  <ul className="slider-lvl1__slider-nav js-slider-hor__slider-nav">
                    <Slider asNavFor={nav1} ref={(slider2) => setNav2(slider2)} {...settings2}>
                      {Array.from(Array(posters.length).keys()).map((index)=>{
                          return <li className="slider-nav-item"><span>{index+1<10?"0"+(index+1):index+1}</span></li>
                      })}
                    </Slider>
                  </ul>
                </div>
                <div className="arr-down" onClick={()=>nav1.slickNext()}>
                  <svg width="14" height="19" viewBox="0 0 14 19" xmlns="http://www.w3.org/2000/svg">
                    <path d="M7 17L6.29289 17.7071L7 18.4142L7.70711 17.7071L7 17ZM12.2929 10.2929L6.29289 16.2929L7.70711 17.7071L13.7071 11.7071L12.2929 10.2929ZM7.70711 16.2929L1.70711 10.2929L0.292893 11.7071L6.29289 17.7071L7.70711 16.2929ZM8 17V0H6V17H8Z" />
                  </svg>
                </div>
              </div>
            </div>
          </div>
          <div className="lvl-row-small">
            <div className="lvl-row-small__img"><img src="img/img2.svg" alt="" /></div>
            <div className="lvl-row-small__text">
              В переводе с итальянского AVANTA (avanti) - вперёд. Дома для жизни и отдыха с высоким уровнем комфорта у моря. Так мы видим ещё один шаг вперёд для резидентов нашего жилого комплекса
            </div>
          </div>
        </div>
      </div>
    </div>
}