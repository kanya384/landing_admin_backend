import React, { Suspense, useEffect, useContext } from "react"
import { SectionsContext } from "../context/sectionsContext"
import { Menu } from '../sections/menu';
import {Properties} from "../sections/properties";
import "../scss/style.scss"
import { Posters } from "../sections/posters";
import { Walks } from "../sections/walks";
import { Excursion } from "../sections/excursion";
import { About } from "../sections/about";
import { Presentation } from "../sections/presentation";
import { Plans } from "../sections/plans";
import { HowToBuy } from "../sections/how_to_buy";
import { Parallax } from "../components/parallax";
import { Video } from "../sections/video";
import { Advantages } from "../sections/advantages";
import { GetPodbor } from "../sections/get_podbor";
import { Infra } from "../sections/infra";
import { HodStr } from "../sections/hod_str";
import { Question } from "../sections/question";
import { Footer } from "../components/footer";


export const Loader = () => {
    //const Location = React.lazy(() => import("./location"))


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
    return generateHtml()

}
