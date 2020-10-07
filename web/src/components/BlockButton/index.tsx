import React from 'react';
import styled, { css } from 'styled-components';

export interface BlockButtonProps {
  label: string;
  size?: 'small' | 'medium' | 'large';
}

export const BlockButton: React.FC<BlockButtonProps> = ({
  label,
  size = 'large',
  ...props
}) => {
  return (
    <BlockButtonStyled size={size} {...props}>{label}</BlockButtonStyled>
  )
}

const BlockButtonStyled = styled.button<BlockButtonProps>`
  border: 2px dashed #607D8B;
  border-radius: 10px;
  padding: 15px 20px 15px 20px;
  width: 370px;
  text-align: center;
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
`