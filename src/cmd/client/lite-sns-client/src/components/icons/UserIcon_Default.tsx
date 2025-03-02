import { css } from "@emotion/css"

type Props = {
    width?: string
    height?: string
    headColor?: string
    bodyColor?: string
    backgroundColor?: string
}

export default function UserIcon_Default(props: Props) {
    const { width, height, headColor, bodyColor, backgroundColor } = props

    const svgStyle = css`
        width: ${width ? width : "256px"};
        height: ${height ? height : "256px"};
    `
    const gStyle = css`
        background-color: ${backgroundColor ? backgroundColor : "#F0F0F0"};
    `
    const headStyle = css`
        fill: ${headColor ? headColor : "#4B4B4B"};
    `
    const bodyStyle = css`
        fill: ${bodyColor ? bodyColor : "#4B4B4B"};
    `

    return (
        // <!--?xml version="1.0" encoding="utf-8"?-->
        // <!-- Generator: Adobe Illustrator 18.0.0, SVG Export Plug-In . SVG Version: 6.00 Build 0)  -->
        <svg className={svgStyle} version="1.1" id="_x32_" xmlns="http://www.w3.org/2000/svg" xmlnsXlink="http://www.w3.org/1999/xlink" x="0px" y="0px" viewBox="0 0 512 512" xmlSpace="preserve">
            <g className={gStyle}>
                <path className={headStyle} d="M255.999,250.486c69.178,0,125.25-56.068,125.25-125.242C381.249,56.072,325.177,0,255.999,0
                    c-69.168,0-125.24,56.072-125.24,125.244C130.759,194.418,186.831,250.486,255.999,250.486z"></path>
                <path className={bodyStyle} d="M319.313,289.033h-63.314h-63.314c-41.289,24.775-77.07,82.58-77.07,132.129c0,30.274,0,90.838,0,90.838
                    h140.385h140.387c0,0,0-60.564,0-90.838C396.386,371.613,360.603,313.808,319.313,289.033z"></path>
            </g>
        </svg>
    )
}
