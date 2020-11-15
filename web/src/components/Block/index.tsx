import React from "react";
import PropTypes from "prop-types";
import styled, { css } from "styled-components";

const BlockStyled = styled.button<BlockProps>`
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
    p.size === "medium" &&
    css`
      width: 177px;
    `}
  ${(p) =>
    p.size === "small" &&
    css`
      width: 112.66px;
    `}
  ${(p) =>
    p.active &&
    css`
      border: none;
    `}
`;

export interface BlockProps {
  children: string;
  active: boolean;
  backgroundColor?: string;
  size?: "small" | "medium" | "large";
}

export const Block: React.FC<BlockProps> = ({
  children,
  active,
  backgroundColor,
  size = "large",
}) => {
  return (
    <BlockStyled size={size} active={active} backgroundColor={backgroundColor}>
      {children}
    </BlockStyled>
  );
};

Block.propTypes = {
  children: PropTypes.string.isRequired,
  active: PropTypes.bool.isRequired,
  backgroundColor: PropTypes.string.isRequired,
  size: PropTypes.oneOf(["small" as const, "medium" as const, "large" as const])
    .isRequired,
};
