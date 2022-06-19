import { SectionsContext } from "./context/sectionsContext"
import { useBlocks } from "./hooks/sections.hook"
import { BrowserRouter as Router, Route, Routes } from "react-router-dom"
import { useSections } from "./hooks/sections.hook";
import Main from "./pages/main";
import { Plans } from "./pages/plans";
import { PlanDetail } from "./pages/plan-detail";
import { Promo } from "./pages/promo";
import { About } from "./pages/about";
import { ContentContext } from "./context/contentContext";
import { useContent } from "./hooks/content.hook";
import { useEffect } from "react";
import { Panel } from "./feature/panel";

function App() {
  const { blocks, setBlocks, menuClick, setMenuClick } = useSections()
  const {content, setContent, administrate, setAdministrate} = useContent()

  useEffect(()=>{
    if (!content) {
      fetch(process.env.REACT_APP_BACKEND+"/content")
        .then((response) => {
          return response.json()
        })
        .then((data) => {
          setContent(data)
        });
    }
  },[])

  return (
    <div className="App container_main">
      <SectionsContext.Provider value={{ blocks, setBlocks, menuClick, setMenuClick }}>
        <ContentContext.Provider value={{ content, setContent, administrate, setAdministrate }}>
        <Router>
          <Panel />
          <div className="over">
              <Routes>
              <Route path="/about" element={<About />} />
                <Route path="/promo" element={<Promo />} />
                <Route path="/plans/:id" element={<PlanDetail />} />
                <Route path="/plans" element={<Plans />} />
                <Route path="*" element={<Main />} />
              </Routes>
            
          </div>
          </Router>
        </ContentContext.Provider>
      </SectionsContext.Provider>
    </div>
  );
}

export default App;