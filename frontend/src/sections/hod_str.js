import { useContext, useEffect, useState } from "react";
import Slider from "react-slick/lib/slider";
import { ContentContext } from "../context/contentContext";

export const HodStr = () => {
  const content = useContext(ContentContext)
  const [nav1, setNav1] = useState();
  const [nav2, setNav2] = useState();
  const [activeYear, setActiveYear] = useState()
  const [months, setMonths] = useState([])
  const [showDropDown, setShow] = useState(false)
  const [activeMonth, setActiveMonth] = useState()
  const [photos, setPhotos] = useState([])

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

  useEffect(()=>{
    if (activeYear === undefined) {
      setActiveYear(content.content?.Years[0])
    }
  },[content.content])

  useEffect(()=>{
    if (activeYear!==undefined){
      let months = []
      content.content.Months.map((month)=>{
        if (month.year_id == activeYear.ID) {
          months.push(month)
        }
      })
      months = months.reverse()
      setMonths(months)
      setActiveMonth(months[0])
    }
  },[activeYear])

  useEffect(()=>{
    if (activeMonth!==undefined){
      let photos = []
      content.content.Photos.map((photo)=>{
        if (photo.month_id === activeMonth?.ID) {
          photos.push(photo.image)
        }
      })
      setPhotos(photos)
    }
  },[activeMonth])

  return (
    <div className="lvl14 parrallax-scene3-end">
      <div className="wrapper">
        <div className="construction-progress">
          <div className="title">Ход строительства</div>
          <div className="construction-progress__filters">
            <div className="filter-left">
              <button className="filter-link active">Фото</button>
              {/*<button className="filter-link">Видео</button>*/}
            </div>
            <div className="filter-right">
              <div className="filter-years">
                {content.content?.Years.map((year)=>{
                  return  <button className={activeYear&& year.ID === activeYear.ID?"filter-link active":"filter-link"}>{year.value}</button>
                })}
              </div>
              
              <div className="dropdown-filter js-dropdown-filter">
                <div className="dropdown-filter__label">
                  <div className="dropdown-filter__label-text" onClick={()=>{setShow(!showDropDown)}}>{activeMonth?.name}</div>
                  <div className="dropdown-filter__arr">
                    <svg width="14" height="9" viewBox="0 0 14 9" xmlns="http://www.w3.org/2000/svg">
                      <path d="M13 1L7 7L1 1" strokeWidth="2"/>
                    </svg>
                  </div>
                </div>
                <div className="dropdown-filter__drop" style={{display:showDropDown?"block":"none"}}>
                  <ul className="dropdown-filter__list">
                    {months.map((month)=>{
                      return  <li onClick={()=>{setShow(false); setActiveMonth(month)}}><span>{month.name}</span></li>
                    })}
                  </ul>
                </div>
              </div>
            </div>
          </div>
          <div className="construction-progress__slider">
            <div className="slider-construction-progress js-slider-construction-progress-wrapper">
              <Slider asNavFor={nav2} ref={(slider1) => setNav1(slider1)} className="slider slider-construction-progress__slider-for js-slider-construction-progress__slider-for" {...settings}>
              
                {photos.map((photo)=>{
                  return  <div className="slide"><img src={process.env.REACT_APP_BACKEND+"store/"+photo} alt=""/></div>
                })}

              </Slider>
              <div className="slider-construction-progress-nav">
                <div className="arr-left" onClick={()=> nav2.slickPrev()}>
                  <svg width="20" height="20" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
                    <path d="M1.5 10.5L0.792893 9.79289L0.0857865 10.5L0.792893 11.2071L1.5 10.5ZM8.20711 15.7929L2.20711 9.79289L0.792893 11.2071L6.79289 17.2071L8.20711 15.7929ZM2.20711 11.2071L8.20711 5.20711L6.79289 3.79289L0.792893 9.79289L2.20711 11.2071ZM1.5 11.5L18.5 11.5V9.5L1.5 9.5L1.5 11.5Z" />
                  </svg>
                </div>
                <div className="nav-build">
                  {photos.length>0?<Slider asNavFor={nav1} ref={(slider2) => setNav2(slider2)} className={"slider-construction-progress__slider-nav js-slider-construction-progress__slider-nav"} {...settings2}>
                    {[...Array(photos.length).keys()].map((val) => {
                      return  <div className="slider-nav-item"><span>{(val+1)<10?"0"+(val+1):(val+1)}</span></div>
                    })}
                  </Slider>:""}
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

export default HodStr
