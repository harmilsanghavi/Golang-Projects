import { css } from "@emotion/react";
import { ClipLoader } from "react-spinners";

const override = css`
  display: block;
  margin: 0 auto;
  border-color: red;
`;

const Spinner = ({ loading }) => {
  return (
    <div className="sweet-loading">
      <ClipLoader css={override} size={150} loading={loading} />
    </div>
  );
};

export default Spinner;