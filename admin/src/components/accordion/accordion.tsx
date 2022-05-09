import { useState } from "react"


export interface AccordionItem {
  title: string,
}

export interface AccordionProps {
  items: AccordionItem[],
}

export const Accordion = (props: AccordionProps) => {
  const [show, setShow] = useState<number[]>([])
  const itemClick = (i:number) => {
    if (!show.includes(i)) {
      setShow([...show, i])
    } else {
      let filtered = show.filter((index) => index!==i)
      setShow(filtered)
    }
  }
  if (!props.items || props.items.length == 0) {
    return <div></div>
  }
  return ( <div className="accordion standard-accordion" id="accordionExample2">
    {props.items.map((item) => {
      return  <div className="accordion-item" onClick={() => itemClick(0)}>
                <h2 className="accordion-header" id="headingFour"><button className="accordion-button collapsed" type="button" data-bs-toggle="collapse" data-bs-target="#collapseFour" aria-expanded="false" aria-controls="collapseFour">{item.title}</button></h2>
                <div className={show.includes(0)?"accordion-collapse collapse show":"accordion-collapse collapse"} id="collapseFour" aria-labelledby="headingFour" data-bs-parent="#accordionExample2">
                  <div className="accordion-body pt-0">
                    <strong>This is the first item's accordion body.</strong> It is shown by default, until the collapse plugin adds the appropriate classes that we use to style each element. These classes control the overall appearance, as well as the showing and hiding via CSS transitions. You can modify any of this with custom CSS or overriding our default variables. It's also worth noting that just about any HTML can go within the <code>.accordion-body</code>, though the transition does limit overflow.
                  </div>
                </div>
              </div>
    })}
    </div>
  )
}