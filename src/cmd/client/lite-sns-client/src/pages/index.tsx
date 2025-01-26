import { css } from "@emotion/css";
import Button from "../components/molecules/button";
import { useNavigate } from "react-router";
import Card from "../components/atoms/card";

export default function Home() {
  const navigate = useNavigate();

  const actionAreaStyle = css`
    display: flex;
    flex-direction: row;
    flex-wrap: wrap;
    justify-content: center;
    align-items: center;
    align-content: center;
  `;

  const signupButtonContainerStyle = css`
    margin: 8px;
  `;

  const signinButtonContainerStyle = css`
    margin: 8px;
  `;

  return (
    <>
      {/* <div>
                Go to <a href="/test">Test Page</a>
            </div> */}
      <div>Lite SNS</div>
      <Card topBorder>
        <div className={actionAreaStyle}>
          <div className={signupButtonContainerStyle}>
            <Button
              onClick={() => {
                navigate("/signup");
              }}
              width="80px"
            >
              Sign up
            </Button>
          </div>
          <div className={signinButtonContainerStyle}>
            <Button
              onClick={() => {
                navigate("/signin");
              }}
              width="80px"
            >
              Sign in
            </Button>
          </div>
        </div>
      </Card>
    </>
  );
}
