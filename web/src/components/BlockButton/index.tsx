import React from 'react';
import styled, { css } from 'styled-components';

export interface BlockButtonProps {
  children: string;
  active: boolean;
  backgroundColor?: string;
  size?: 'small' | 'medium' | 'large';
}

export const BlockButton: React.FC<BlockButtonProps> = ({
  children,
  active,
  backgroundColor,
  size = 'large',
  ...props
}) => {
  return (
    <BlockButtonStyled
      size={size}
      active={active}
      backgroundColor={backgroundColor}
      {...props}
    >{children}</BlockButtonStyled>
  )
}

const BlockButtonStyled = styled.button<BlockButtonProps>`
  border: 2px dashed #607D8B;
  border-radius: 10px;
  padding: 15px 20px 15px 20px;
  width: 370px;
  text-align: center;
  background-color: ${p => p.backgroundColor};
  :hover{
    cursor: pointer;
  }
  :active{
    background: #2ab8b96b;
  }
  :focus{
    outline: 0;
  }
  ${p => p.size === 'medium' && css`
    width: 177px;
  `}
  ${p => p.size === 'small' && css`
    width: 112.66px;
  `}
  ${p => p.active && css`
    border: none;
  `}
`