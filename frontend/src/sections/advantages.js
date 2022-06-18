import { SliderWithText } from "../components/slider-with-text";
import { ContentContext } from "../context/contentContext";
import { useContext, useEffect, useState } from "react";

export const Advantages = () => {
  const content = useContext(ContentContext)
  const [sliders, setSliders] = useState()

  useEffect(()=>{
    if (content.content) {
      let sliders = []
      content.content.Advantages.map((advantage) => {
        let slides = [];
        content.content.AdvantagePhotos.map((photo) =>{
          if (photo.AdvantageID == advantage.ID) slides.push(process.env.REACT_APP_BACKEND+"store/"+ photo.image)
        })
        sliders.push({
          title: advantage.title,
          text: advantage.description,
          slides: slides,
        })
      })
      setSliders(sliders)
    }
  },[content.content])

  return( <div className="lvl11">
            <div className="wrapper">
              <h2 className="lvl-title">А также вас ждёт:</h2>
              <div className="about-info-list">
                {sliders&& sliders.map((slider, index)=>{
                  return <SliderWithText slides={slider} key={"slider"+index} />
                })}
              </div>  
            </div>
          </div>
  )
}