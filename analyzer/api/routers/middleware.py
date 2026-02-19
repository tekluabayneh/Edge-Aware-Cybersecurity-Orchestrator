# async def getUserJwt(req:Request, call_next):
#     header = req.header("Authorization", "Bearer "+token)
#     if not header: 
#         print("no header found")
#         return 
#
#     token = header.split("")[1]
#
from fastapi import Request, status
from starlette.middleware.base import BaseHTTPMiddleware
from starlette.responses import JSONResponse, Response

class JWTAuthMiddleware(BaseHTTPMiddleware):
    async def dispatch(self, request: Request, call_next):
        excluded_paths = {
            "/docs",
        }

        if request.method == "OPTIONS" or request.url.path in excluded_paths:
            return await call_next(request)

        # Get Authorization header
        auth_header = request.headers.get("Authorization")
        if not auth_header:
            return JSONResponse(
                status_code=status.HTTP_401_UNAUTHORIZED,
                content={"detail": "Not authenticated - missing Authorization header"},
                headers={"WWW-Authenticate": "Bearer"},
            )

        # Parse Bearer token
        scheme, _, token = auth_header.partition(" ")
        if scheme.lower() != "bearer" or not token.strip():
            return JSONResponse(
                status_code=status.HTTP_401_UNAUTHORIZED,
                content={"detail": "Invalid authentication scheme or empty token"},
                headers={"WWW-Authenticate": "Bearer"},
            )

        try:
          pass 

        except Exception as e:
          return JSONResponse(
            status_code=401,
                content={"detail": f"opration failed: {str(e)}"}
            )

        response: Response = await call_next(request)
        return response   
    # store it in wehn ever any kiinds of api called enclude the curretn token



