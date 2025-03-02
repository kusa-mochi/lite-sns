import { css } from "@emotion/css";
import { ReactNode, useState } from "react"
import MenuBar from "../organisms/menu_bar";

type Props = {
    children?: ReactNode
}

export default function UserPage(props: Props) {
    const { children } = props
    const [menuHeight, setMenuHeight] = useState(48)

    const pageStyle = css`
        width: 100%;
        height: 100%;
        overflow-x: hidden;
        overflow-y: auto;

        scrollbar-width: none;
        -ms-overflow-style: none;

        &::webkit-scrollbar {
        display: none;
        }
    `;
    const timelineEndStyle = css`
        width: 100%;
        margin: ${menuHeight}px 0;
    `;

    return (
        <div className={pageStyle}>
            {children}
            <div className={timelineEndStyle}></div>
            <MenuBar height={`${menuHeight}px`} />
        </div>
    )
}
