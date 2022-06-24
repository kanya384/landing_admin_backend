import React, { Suspense, useEffect, useContext } from "react"
import { SectionsContext } from "../context/sectionsContext"
import { Menu } from '../sections/menu';
import "../scss/style.scss"
import { Posters } from "../sections/posters";
import { About } from "../sections/about";
import { Parallax } from "../components/parallax";
import { ContentContext } from "../context/contentContext";


export const Loader = () => {
    const Properties = React.lazy(() => import("../sections/properties"))
    const Walks = React.lazy(() => import("../sections/walks"))
    const Excursion = React.lazy(() => import("../sections/excursion"))
    const Presentation = React.lazy(() => import("../sections/presentation"))
    const Plans = React.lazy(() => import("../sections/plans"))
    const HowToBuy = React.lazy(() => import("../sections/how_to_buy"))
    const Video = React.lazy(() => import("../sections/video"))
    const Advantages = React.lazy(() => import("../sections/advantages"))
    const GetPodbor = React.lazy(() => import("../sections/get_podbor"))
    const Infra = React.lazy(() => import("../sections/infra"))
    const HodStr = React.lazy(() => import("../sections/hod_str"))
    const Question = React.lazy(() => import("../sections/question"))
    const Footer = React.lazy(() => import("../components/footer"))


    const blocksImports = [
        <Menu />,
        <Posters />,
        <Parallax image={"img/title-bg1.svg"} />,
        <Properties />,
        <Walks />,
        <Excursion />,
        <About />,
        <Presentation />,
        <Plans />,
        <HowToBuy />,
        <Parallax image={"img/title-bg2.svg"} />,
        <Video />,
        <Advantages />,
        <GetPodbor />,
        <Infra />,
        <Parallax image={"img/title-bg3.svg"} />,
        <HodStr />,
        <Question />,
        <Footer />
    ]

    const loaded = useContext(SectionsContext)
    const content = useContext(ContentContext)

    const LoadBlock = (block) => {
        return <Suspense fallback={<div>Загрузка...</div>}>{block}</Suspense>
    }

    const generateHtml = () => {
        let toDraw = []
        for (let i = 0; i < loaded.blocks; i++) {
            toDraw.push(LoadBlock(blocksImports[i]))
        }
        return (

            <div className="blocks" data={loaded.menuClick ? "true" : ""}>
                {toDraw.map((block) =>
                    block
                )}
            </div>

        )
    }
    const handleScroll = (event) => {
        if (loaded.blocks < blocksImports.length) {
            loaded.setBlocks(blocksImports.length)
        } else {
            window.removeEventListener('scroll', handleScroll, true);
        }
    }
    useEffect(() => {
        if (loaded.blocks < blocksImports.length) {
            window.addEventListener('scroll', handleScroll, true);
        }
    })

    useEffect(()=>{

        if (content.administrate && (loaded.blocks < blocksImports.length)) {
            let y = window.scrollY
            loaded.setBlocks(blocksImports.length)
            setTimeout(()=>{
                window.scrollTo(0, y)
            },400)
        }
    },[content.administrate])
    return generateHtml()

}
