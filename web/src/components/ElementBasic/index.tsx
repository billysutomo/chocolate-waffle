import React from "react";
import styled, { css } from "styled-components";

const ElementStyled = styled.button<ElementBasicProps>`
  border: 2px dashed ;
  border-color: rgba(255,255,255,0.5);
  border-radius: 10px;
  padding: 15px 20px 15px 20px;
  color: rgba(255,255,255,0.5);
  font-weight: bolder;
  font-size: 16px;
  width: 370px;
  text-align: center;
  margin-bottom: 16px;
  background-color: transparent;
  margin: 0 auto;
  margin-left: 8px;
  margin-right: 8px;
  margin-bottom: 16px;
  :hover {
    cursor: pointer;
  }
  :active {
    background: #2ab8b96b;
  }
  :focus {
    outline: 0;
  }
`;

export interface ElementBasicProps {
  onClick?: React.MouseEventHandler<HTMLButtonElement>;
}

export const ElementBasic: React.FC<ElementBasicProps> = ({
  children,
  onClick
}) => {


  return (
    <ElementStyled onClick={onClick}>
      {children}
    </ElementStyled>
  );
};
