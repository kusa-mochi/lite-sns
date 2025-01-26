import { ReactNode } from "react";
import { useTheme } from "../../providers/themeProvider";
import { css } from "@emotion/css";

type Props = {
  children?: ReactNode;
  isOpen?: boolean;
  onClickBackdrop?: () => void;
};

export default function Modal(props: Props) {
  const { children, isOpen, onClickBackdrop } = props;
  const theme = useTheme();

  const modalRootStyle = css`
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;

    display: flex;
    flex-direction: column;
    flex-wrap: nowrap;
    justify-content: center;
    align-items: center;

    transition: opacity 0.2s linear;
  `;

  const backdropStyle = css`
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background-color: rgba(0, 0, 0, 0.3);
  `;

  const modalStyle = css`
    position: relative;
    width: 90%;
    max-height: 90%;
    padding: 2rem;
    border-radius: 1rem;
    background-color: ${theme.palette.base.bodyBackgroundColor};
    box-shadow: rgba(0, 0, 0, 0.3) 0px 3px 4px 0px;
  `;

  return (
    isOpen && (
      <div className={modalRootStyle}>
        <div
          className={backdropStyle}
          onClick={() =>
            onClickBackdrop === undefined ? {} : onClickBackdrop()
          }
        ></div>
        <div className={modalStyle}>{children}</div>
      </div>
    )
  );
}
