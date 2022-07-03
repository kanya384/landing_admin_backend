import { SectionsContext } from "./context/sectionsContext"
import { BrowserRouter as Router, Route, Routes } from "react-router-dom"
import { useSections } from "./hooks/sections.hook";
import Main from "./pages/main";
import { ContentContext } from "./context/contentContext";
import { useContent } from "./hooks/content.hook";
import React, { Suspense, useEffect } from "react";
import { Panel } from "./feature/panel";

function App() {
  const { blocks, setBlocks, menuClick, setMenuClick } = useSections()
  const {content, setContent, administrate, setAdministrate} = useContent()
  const PlanDetail = React.lazy(() => import("./pages/plan-detail"))
  const Plans = React.lazy(() => import("./pages/plans"))
  const Promo = React.lazy(() => import("./pages/promo"))
  const About = React.lazy(() => import("./pages/about"))

  useEffect(()=>{
    if (!content) {
      if (typeof contentTemplate !== 'undefined') {
        setContent(contentTemplate)
      } else {
        fetch(process.env.REACT_APP_BACKEND+"/content")
          .then((response) => {
            return response.json()
          })
          .then((data) => {
            setContent(data)
          });
        }
      }
  },[])

  return (
    <div className="App container_main">
      <ContentContext.Provider value={{ content, setContent, administrate, setAdministrate }}>
      <SectionsContext.Provider value={{ blocks, setBlocks, menuClick, setMenuClick }}>
        
        <Router>
          <Panel />
          <div className="over">
              <Routes>
              <Route path="/about" element={<Suspense  fallback={<div>Загрузка...</div>}><About /></Suspense>} />
                <Route path="/promo" element={<Suspense  fallback={<div>Загрузка...</div>}><Promo /></Suspense>} />
                <Route path="/plans/:id" element={<Suspense  fallback={<div>Загрузка...</div>}><PlanDetail /></Suspense>} />
                <Route path="/plans" element={<Suspense  fallback={<div>Загрузка...</div>}><Plans /></Suspense>} />
                <Route path="*" element={<Suspense  fallback={<div>Загрузка...</div>}><Main /></Suspense>} />
              </Routes>
          </div>
          </Router>
        
      </SectionsContext.Provider>
      </ContentContext.Provider>
    </div>
  );
}

export default App;