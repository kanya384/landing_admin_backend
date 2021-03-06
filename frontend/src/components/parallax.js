import { useContext, useEffect, useRef, useState } from "react"
import { useIsVisible } from 'react-is-visible'
import { SectionsContext } from "../context/sectionsContext"

export const Parallax = ({image}) => {
  const loaded = useContext(SectionsContext)
  const [id, setId] = useState(image.replaceAll('/', '').replaceAll('.', ''))
  const nodeRef = useRef()
  const isVisible = useIsVisible(nodeRef)
  const [translate, setTranslate] = useState(2500)

  useEffect(()=>{
    if (isVisible) {
      setTranslate(-2500)
    } else {
      setTranslate(3000)
    }
  },[isVisible])

  return (<section ref={nodeRef} id={id} className={"js-parrallax parrallax-scene1 "+id}>
            <div className="fake-img fake-img1" ><img src={loaded.blocks>3?image:""} alt="" /></div>
            <div key={id} className={translate==3000?"":"parralax-img1"} style={{position: "absolute", transform: `translateX(${translate}px)`}}><img src={loaded.blocks>3?image:""} alt="" /></div>
          </section>
  )
}