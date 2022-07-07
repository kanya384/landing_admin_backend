import React, { useContext } from "react"
import { SectionsContext } from "../context/sectionsContext"
import {Loader} from "../feature/loader"
import PlanDetail from "./plan-detail"

export const Main = () => {
  const sections = useContext(SectionsContext)
  return <React.Fragment>
          <div style={{display:sections.showPlan?"none":""}}>
            <Loader />
          </div>
          <div style={{display:sections.showPlan?"":"none"}}>
            <PlanDetail ident={sections.showPlan} />
          </div>

          
        </React.Fragment>
}

export default Main