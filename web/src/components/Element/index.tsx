import React from "react";
import styled, { css } from "styled-components";

const ElementStyled = styled.button<ElementProps>`
  border: 2px dashed #607d8b;
  border-radius: 10px;
  padding: 15px 20px 15px 20px;
  width: 370px;
  text-align: center;
  margin-bottom: 16px;
  background-color: ${(p) => p.backgroundColor};
  :hover {
    cursor: pointer;
  }
  :active {
    background: #2ab8b96b;
  }
  :focus {
    outline: 0;
  }
  ${(p) =>
    p.size === Sizes.medium &&
    css`
      width: 177px;
    `}
  ${(p) =>
    p.size === Sizes.small &&
    css`
      width: 112.66px;
    `}
  ${(p) =>
    p.active &&
    css`
      border: none;
    `}
`;

export enum Sizes {
  small,
  medium,
  large,
}

export enum Types {
  messenger,
  block,
  social_link,
  nothing,
}

export interface ElementProps {
  children: string;
  active: boolean;
  backgroundColor?: string;
  size?: Sizes;
  elementType: Types;
}

export const Element: React.FC<ElementProps> = ({
  children,
  active,
  backgroundColor,
  size,
  elementType,
}) => {
  return (
    <ElementStyled
      size={size}
      active={active}
      backgroundColor={backgroundColor}
      elementType={elementType}
    >
      {children}
    </ElementStyled>
  );
};
