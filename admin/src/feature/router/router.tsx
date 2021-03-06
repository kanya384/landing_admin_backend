import { useTypedSelector } from "../../hooks/use-typed-selector"
import { Authentication } from "../../pages/authentication"
import { BrowserRouter, Route, Routes } from "react-router-dom"
import Spinner from "../../components/spinner/spinner"
import { useActions } from "../../hooks/use-actions"
import React, { useEffect } from "react"
import MenuComponent from "../menu-component"
import NavbarComponent from "../../components/navbar-component"
import { Posters } from "../../pages/posters"
import { Hod } from "../../pages/hod"
import { Advantages } from "../../pages/advantages"
import { AdvantageDetail } from "../../pages/advantage_detail"
import { Plans } from "../../pages/plans"
import { Doc } from "../../pages/docs"
import { Videos } from "../../pages/videos"
import { Leads } from "../../pages/leads"
import { ProjectInfos } from "../../pages/projectInfos"
import { Titles } from "../../pages/titles"
import { Settings } from "../../pages/settings"
import { Actions } from "../../pages/actions"

const Router: React.FC = () => {
  const auths = useTypedSelector(({ auths }) => {
    return auths
  })
  const { checkAuth } = useActions();

  useEffect(()=>{
    checkAuth()
  }, [checkAuth])

  if (auths?.loading === true && auths?.auth === null) {
    return <Spinner />
  } else if (auths?.auth === true) {
    return (
      <main className="main">
        <div className="container-fluid px-0">
        <BrowserRouter basename="/admin">
          <MenuComponent />
          <NavbarComponent />
            <div className="content pt-5">
              <Routes>
                <Route path="/" element={<Leads />} />
                <Route path="/project-info" element={<ProjectInfos />} />
                <Route path="/actions" element={<Actions />} />
                <Route path="/settings" element={<Settings />} />
                <Route path="/titles" element={<Titles />} />
                <Route path="/video" element={<Videos />} />
                <Route path="/docs" element={<Doc />} />
                <Route path="/plans" element={<Plans />} />
                <Route path="/advantages/:id" element={<AdvantageDetail />} />
                <Route path="/advantages" element={<Advantages />} />
                <Route path="/progress" element={<Hod />} />
                <Route path="/posters" element={<Posters />} />
                <Route path="*" element={"good"} />
              </Routes>
            </div>
          </BrowserRouter>
        </div>
      </main>
    )
  } else {
    console.log(auths)
    return (
      <BrowserRouter>
        <Routes>
          <Route path="*" element={<Authentication />} />
        </Routes>
      </BrowserRouter>
    )
  }
}

export default Router