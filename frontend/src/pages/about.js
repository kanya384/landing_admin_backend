import React, { useState } from "react"
import Slider from "react-slick/lib/slider"
import { AppBar } from "../components/appbar"
import { EditableText } from "../components/editable-text"
import { Footer } from "../components/footer"
import { Modal } from "../components/modals"
import { Form } from "../components/form";

export const About = () => {

  const [nav1, setNav1] = useState();
  const [nav2, setNav2] = useState();
  const [isOpen, setIsOpen] = useState(false)

  let settings = {
    slidesToShow: 3,
    slidesToScroll: 1,
    dots: false,
    variableWidth: true,
    centerMode: true,
    arrows: false,
    responsive: [
        {
            breakpoint: 769,
            settings: {
                slidesToShow: 1,
                slidesToScroll: 1,
                centerMode: false,
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
  /*
  $('.js-slider-about__slider-for').slick({
        asNavFor: $('.js-slider-about-wrapper .js-slider-about__slider-nav'),
       
    });
    $('.js-slider-about__slider-nav').slick({
       
        
    });
  
  */
  return (
    <React.Fragment>
      <AppBar />
      <div class="container">
        <div class="about-banner">
          <div class="b-text">
            <div class="about-banner__title">ava group</div>
            <div class="about-banner__text"><EditableText id={"62aefdb322de07ce5e31751f"} defaultText={"Захватывающая история.<br />Вдохновляющее будущее"}/></div>
          </div>
          <div class="img"><img src="img/about-banner.png" alt="" /></div>
        </div>
        <div class="wrapper">
          <div class="options-room options-room--about">
            <div class="option-room__description">
              <div class="text">
                <EditableText id={"62aefdb322de07ce5e317520"} defaultText={"Девелопмент — это развитие. Планируя каждый проект, мы думаем о том, как будут жить, что чувствовать их жители. Мы думаем о создании цельной среды, в которой не нужно тратить лишнего времени на повседневные хлопоты и можно дарить его близким, думать о комфорте, здоровье и детях, безопасности и экологии. Мы заботимся о развитии — росте благополучия и реализации стремлений наших клиентов."}/>
              </div>
              <div class="b-download">
                <a href="#" class="link-download">
                    <svg class="svg"><use xlinkHref="img/sprite.svg#circle-download"></use></svg> Пректная декларация<br />ЖК AVAnta
                </a>
              </div>
            </div>
            <div class="options-room--about__wrapper">
              <div class="option-room__item-wrap">
                <div class="option-room__item">
                  <div class="label"><EditableText id={"62aefdb322de07ce5e317521"} defaultText={"Завершенных проектов"}/></div>
                  <div class="value"><EditableText id={"62aefdb322de07ce5e317522"} defaultText={"100+"}/></div>
                </div>
              </div>
              <div class="option-room__item-wrap">
                <div class="option-room__item">
                  <div class="label"><EditableText id={"62aefdb322de07ce5e317523"} defaultText={"Семей обрели новый дом"}/></div>
                  <div class="value"><EditableText id={"62aefdb322de07ce5e317524"} defaultText={"5000"}/></div>
                </div>
              </div>
              <div class="option-room__item-wrap">
                <div class="option-room__item">
                  <div class="label"><EditableText id={"62aefdb322de07ce5e317525"} defaultText={"Детских садов и школ построено"}/></div>
                  <div class="value"><EditableText id={"62aefdb322de07ce5e317526"} defaultText={"42"}/></div>
                </div>
              </div>
              <div class="option-room__item-wrap">
                <div class="option-room__item">
                  <div class="label"><EditableText id={"62aefdb322de07ce5e317527"} defaultText={"Квадратных метров жилой недвижимости построено"}/></div>
                  <div class="value"><EditableText id={"62aefdb322de07ce5e317528"} defaultText={"1,2 млн"}/></div>
                </div>
              </div>
              <div class="option-room__item-wrap">
                <div class="option-room__item">
                  <div class="label"><EditableText id={"62aefdb322de07ce5e31752a"} defaultText={"Лет на рынке недвижимости"}/></div>
                  <div class="value"><EditableText id={"62aefdb322de07ce5e317529"} defaultText={"15+"}/></div>
                </div>
              </div>
              <div class="option-room__item-wrap">
                <div class="option-room__item">
                  <div class="label"><EditableText id={"62aefdb322de07ce5e31752b"} defaultText={"Сотрудников"}/></div>
                  <div class="value"><EditableText id={"62aefdb322de07ce5e31752c"} defaultText={"1500+"}/></div>
                </div>
              </div>
            </div>
          </div>
          <div class="about-slider">
            <div class="slider">
              <div class="construction-progress__slider">
                <div class="slider-construction-progress js-slider-about-wrapper">
                  <div class="slider-construction-progress-nav--wrapper">
                    <div class="slider-construction-progress-nav">
                      <div class="arr-left" onClick={()=>nav1.slickPrev()}>
                        <svg width="20" height="20" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
                          <path
                            d="M1.5 10.5L0.792893 9.79289L0.0857865 10.5L0.792893 11.2071L1.5 10.5ZM8.20711 15.7929L2.20711 9.79289L0.792893 11.2071L6.79289 17.2071L8.20711 15.7929ZM2.20711 11.2071L8.20711 5.20711L6.79289 3.79289L0.792893 9.79289L2.20711 11.2071ZM1.5 11.5L18.5 11.5V9.5L1.5 9.5L1.5 11.5Z"
                          />
                        </svg>
                      </div>
                      <div class="nav-build">
                        <Slider asNavFor={nav1} ref={(slider2) => setNav2(slider2)} className="slider-construction-progress__slider-nav js-slider-about__slider-nav" {...settings2}>
                          <div class="slider-nav-item"><span>01</span></div>
                          <div class="slider-nav-item"><span>02</span></div>
                          <div class="slider-nav-item"><span>03</span></div>
                          <div class="slider-nav-item"><span>04</span></div>
                        </Slider>
                      </div>
                      <div class="arr-right" onClick={()=>nav1.slickNext()}>
                        <svg width="20" height="20" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
                          <path
                            d="M18.5 9.5L19.2071 10.2071L19.9142 9.5L19.2071 8.79289L18.5 9.5ZM11.7929 4.20711L17.7929 10.2071L19.2071 8.79289L13.2071 2.79289L11.7929 4.20711ZM17.7929 8.79289L11.7929 14.7929L13.2071 16.2071L19.2071 10.2071L17.7929 8.79289ZM18.5 8.5L1.5 8.5L1.5 10.5L18.5 10.5V8.5Z"
                          />
                        </svg>
                      </div>
                    </div>
                  </div>
                  <Slider asNavFor={nav2} ref={(slider1) => setNav1(slider1)} className="slider slider-construction-progress__slider-for js-slider-about__slider-for" {...settings}>
                    <div class="slide">
                      <div class="img"><img src="img/about-slide.png" alt="" /></div>
                      <div class="description">
                        <div class="num">1</div>
                        <div class="text">
                          <EditableText id={"62aefdb322de07ce5e31752d"} defaultText={"Тщательно оцениваем земельный участок, его экологические характеристики, а также действующую жилую, социальную, инженерную и транспортную инфраструктуру с учётом потенциала их дальнейшего развития."}/>
                        </div>
                      </div>
                    </div>
                    <div class="slide">
                      <div class="img"><img src="img/about-slide.png" alt="" /></div>
                      <div class="description">
                        <div class="num">2</div>
                        <div class="text">
                          <EditableText id={"62aefdb322de07ce5e31752e"} defaultText={"Тщательно оцениваем земельный участок, его экологические характеристики, а также действующую жилую, социальную, инженерную и транспортную инфраструктуру с учётом потенциала их дальнейшего развития."}/>
                        </div>
                      </div>
                    </div>
                    <div class="slide">
                      <div class="img"><img src="img/about-slide.png" alt="" /></div>
                      <div class="description">
                        <div class="num">3</div>
                        <div class="text">
                          <EditableText id={"62aefdb322de07ce5e31752f"} defaultText={"Тщательно оцениваем земельный участок, его экологические характеристики, а также действующую жилую, социальную, инженерную и транспортную инфраструктуру с учётом потенциала их дальнейшего развития."}/>
                        </div>
                      </div>
                    </div>
                    <div class="slide">
                      <div class="img"><img src="img/about-slide.png" alt="" /></div>
                      <div class="description">
                        <div class="num">4</div>
                        <div class="text">
                          <EditableText id={"62aefdb322de07ce5e317530"} defaultText={"Тщательно оцениваем земельный участок, его экологические характеристики, а также действующую жилую, социальную, инженерную и транспортную инфраструктуру с учётом потенциала их дальнейшего развития."}/>
                        </div>
                      </div>
                    </div>
                  </Slider>
                </div>
              </div>
            </div>
          </div>
          <div class="about-map">
            <div class="about-map-title">
              <EditableText id={"62aefdb322de07ce5e317531"} defaultText={"AVA Group — девелопер федерального уровня, реализует свои проекты в 4 городах России: Краснодар, Сочи, Анапа, Москва"}/>
            </div>
            <div class="map-wrapper">
              <div class="map-img"><img src="img/big-map.svg" alt="" /></div>
            </div>
          </div>
          <div class="about-form">
            <div class="form-ec">
              <div class="form-ec__title"><EditableText id={"62aefdb322de07ce5e317532"} defaultText={"Если у вас остались вопросы, наши менеджеры вам помогут"}/></div>
              <div class="form-ec__content">
                  <Form fields={[
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
                    description={`Задать вопрос менеджеру`}
                    celtype={"getQuestionAbout"}
                    close={()=>{}} 
                    callback={()=>{setIsOpen(true)}}
                />
              </div>
            </div>
          </div>
        </div>
      </div>
      <Footer />
      <Modal 
          success={true}
          position={window.pageYOffset}
          opened={isOpen}
          close = {()=>{setIsOpen(null)}}
        />
    </React.Fragment>
  )
}