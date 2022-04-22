const DEFAULT_COLOR = {
   red: '#000000',
   yellow: '#000000',
   green: '#000000'
};

export const BEGIN_STATE = {
   umlImgName: 'begin',
   color: DEFAULT_COLOR
};


export const END_STATE = {
   umlImgName: 'end',
   color: DEFAULT_COLOR
};

export const WORKING_STATE = {
   red: {
      umlImgName: 'red',
      color: {
         red: '#cc3232',
         yellow: '#000000',
         green: '#000000'
      }
   },
   green: {
      umlImgName: 'green',
      color: {
         red: '#000000',
         yellow: '#000000',
         green: '#2dc937'
      }
   },
   yellow: {
      umlImgName: 'yellow',
      color: {
         red: '#000000',
         yellow: '#e7b416',
         green: '#000000'
      }
   }
};

export const SYSTEM_ERROR_STATE = {
   default: {
      umlImgName: 'flashingRed',
      color: DEFAULT_COLOR
   },
   red: {
      umlImgName: 'flashingRed',
      color: {
         red: '#cc3232',
         yellow: '#000000',
         green: '#000000'
      }
   }
};


export const STATES = {
   'INITAL_STATE': 'initial',
   'BEGIN_STATE': 'begin',
   'WORKING_STATE': 'working',
   'ERROR_STATE': 'error',
   'END_STATE': 'end'
};