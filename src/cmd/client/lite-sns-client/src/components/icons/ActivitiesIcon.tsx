import { css } from "@emotion/css"

type Props = {
    width?: string
    height?: string
    bodyColor?: string
    backgroundColor?: string
}

export default function ActivitiesIcon(props: Props) {
    const { width, height, bodyColor, backgroundColor } = props

    const svgStyle = css`
        width: ${width ? width : "256px"};
        height: ${height ? height : "256px"};
    `
    const gStyle = css`
        background-color: ${backgroundColor ? backgroundColor : "#F0F0F0"};
    `
    const bodyStyle = css`
        fill: ${bodyColor ? bodyColor : "#4B4B4B"};
    `

    return (
        <svg className={svgStyle} version="1.1" id="_x32_" xmlns="http://www.w3.org/2000/svg" xmlnsXlink="http://www.w3.org/1999/xlink" x="0px" y="0px" viewBox="0 0 512 512" xmlSpace="preserve">
            <g className={gStyle}>
                <path className={bodyStyle} d="M256,0C114.84,0,0,114.844,0,256c0,141.164,114.84,256,256,256s256-114.836,256-256
                    C512,114.844,397.16,0,256,0z M256,451.047c-107.547,0-195.047-87.492-195.047-195.047c0-107.547,87.5-195.047,195.047-195.047
                    S451.047,148.453,451.047,256C451.047,363.555,363.547,451.047,256,451.047z"></path>
                <path className={bodyStyle} d="M258.434,115.758c-12.81,0-23.195,10.383-23.195,23.195v105.008l-74.047,74.047
                    c-9.061,9.054-9.061,23.742,0,32.804c9.058,9.055,23.744,9.055,32.804,0l87.635-87.633v-23.766V138.953
                    C281.631,126.141,271.246,115.758,258.434,115.758z"></path>
            </g>
        </svg>
    )
}
