import { FC } from "react";
import { Plan } from "../../api";


export const PlanCard: FC<{plan: Plan, btnClick: any}> = (props) => {
  const flatTitle = (): string => {
    switch (props.plan.rooms) {
      case 1:
         return "Однокомнатная квартира"
      case 2:
        return "Двухкомнатная квартира"
      case 3:
        return "Трехкомнатная квартира"
      default:
        return "Квартира-студия"
    }
  }
  return( <div className='col-sm-6 col-md-4 col-lg-3'>
            <div className="card">
              <img style={{display:"flex", justifyContent: "center"}} className="card-img-top" src={process.env.REACT_APP_BACKEND_URL!+"/store/"+props.plan.image} alt="нет фото" />
              <div className="card-body">
                <button className="btn btn-primary col-12 mb-2" onClick={()=>{props.btnClick(props.plan.id)}}>Изменить фото</button>
                <h5 className="card-title mb-2">{flatTitle()}</h5>  
                <p className="card-text">
                  <div className="d-flex justify-content-between">
                    <span>Статус</span>
                    <strong>{props.plan.status===true?"Свободна":"Снята с продажи"}</strong>
                  </div>
                  <div className="d-flex justify-content-between">
                    <span>Стоимость</span>
                    <strong>{(Math.round(props.plan.price!/10000)/100)} млн руб.</strong>
                  </div>
                  <div className="d-flex justify-content-between">
                    <span>Общая площадь</span>
                    <span>{props.plan.area}м<sup>2</sup></span>
                  </div>
                  <div className="d-flex justify-content-between">
                    <span>Жилая площадь</span>
                    <span>{props.plan.living_area}м<sup>2</sup></span>
                  </div>
                </p>
                
              </div>
            </div>
          </div>
  )

}