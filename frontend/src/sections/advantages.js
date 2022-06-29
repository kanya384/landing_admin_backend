import { SliderWithText } from "../components/slider-with-text";
import { ContentContext } from "../context/contentContext";
import React, { useContext, useEffect, useState } from "react";
import { Link } from "react-router-dom";
import { EditableText } from "../components/editable-text";

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

  return( <React.Fragment>
            <div className="lvl11">
              <div className="wrapper">
                <h2 className="lvl-title">А также вас ждёт:</h2>
                <div className="about-info-list">
                  {sliders&& sliders.map((slider, index)=>{
                    return <SliderWithText slides={slider} key={"slider"+index} />
                  })}
                </div>  
              </div>
            </div>
            <div className="wrapper">
          <div className="b-link-row">
              <Link to="/promo" className="lnk-params js-open-modal-flat">
                <span className="lnk-params__ico">
                <svg className="ico" width="66" height="30" viewBox="0 0 66 30"
                      xmlns="http://www.w3.org/2000/svg">
                  <path d="M64.5 14.9983L65.2033 15.7092L65.9219 14.9982L65.2032 14.2874L64.5 14.9983ZM51.0509 29.7109L65.2033 15.7092L63.7967 14.2875L49.6443 28.2891L51.0509 29.7109ZM65.2032 14.2874L51.0508 0.289034L49.6444 1.71097L63.7968 15.7093L65.2032 14.2874ZM64.5 13.9983H0.5V15.9983H64.5V13.9983Z" />
                </svg>
                </span>
                <EditableText id={"62af012c2f3ab6f9b4a854dc"} defaultText={"И еще множество причин жить на юге"}/>
              </Link>
          </div>
        </div>
      </React.Fragment>
  )
}

export default Advantages